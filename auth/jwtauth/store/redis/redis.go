package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

type Config struct {
	Address   string
	DB        int
	Password  string
	KeyPrefix string
}

type redisClienter interface {
	Get(key string) *redis.StringCmd
	Set(key string, value interface{}, expiration time.Duration) *redis.StatusCmd
	Expire(key string, expiration time.Duration) *redis.BoolCmd
	Exists(keys ...string) *redis.IntCmd
	TxPipeline() redis.Pipeliner
	Del(keys ...string) *redis.IntCmd
	Close() error
}

type Store struct {
	cli    redisClienter
	prefix string
}

func (store *Store) NewStore(cfg *Config) *Store {
	cli := redis.NewClient(&redis.Options{
		Addr:     cfg.Address,
		DB:       cfg.DB,
		Password: cfg.Password,
	})
	return &Store{
		cli:    cli,
		prefix: cfg.KeyPrefix,
	}
}
func NewStoreWithClient(cli *redis.Client, KeyPrefix string) *Store {
	return &Store{
		cli:    cli,
		prefix: KeyPrefix,
	}
}

func NewStoreWithClusterClient(cli *redis.ClusterClient, KeyPrefix string) *Store {
	return &Store{
		cli:    cli,
		prefix: KeyPrefix,
	}
}

func (s *Store) wrapperKey(key string) string {
	return fmt.Sprintf("%s%s", s.prefix, key)
}

func (s *Store) Set(ctx context.Context, tokenString string, expiration time.Duration) error {
	cmd := s.cli.Set(s.wrapperKey(tokenString), 1, expiration)
	return cmd.Err()
}
func (s *Store) Check(ctx context.Context, tokenString string) (bool, error) {
	cmd := s.cli.Exists(s.wrapperKey(tokenString))
	if err := cmd.Err(); err != nil {
		return false, err
	}
	return cmd.Val() > 0, nil
}

func (s *Store) Delete(ctx context.Context, tokenString string) (bool, error) {
	cmd := s.cli.Del(s.wrapperKey(tokenString))
	if err := cmd.Err(); err != nil {
		return false, err
	}
	return cmd.Val() > 0, nil
}
func (s *Store) Close() error {
	return s.cli.Close()
}