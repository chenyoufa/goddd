package test

import (
	"context"
	"fmt"
	"gdemo1/auth/jwtauth/store/redis"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	addr = "127.0.0.1:6379"
)

func TestStore(t *testing.T) {

	fmt.Println()

	store := redis.NewStore(&redis.Config{
		Addr:      addr,
		DB:        1,
		KeyPrefix: "prefix",
	})
	defer store.Close()
	key := "test"
	ctx := context.Background()
	err := store.Set(ctx, key, 0)
	assert.Nil(t, err)

	b, err := store.Check(ctx, key)
	assert.Nil(t, err)
	assert.Equal(t, true, b)

	// b, err = store.Delete(ctx, key)
	// assert.Nil(t, err)
	// assert.Equal(t, false, b)
}
