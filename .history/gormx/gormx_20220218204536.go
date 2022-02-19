package gormx

import (
	"database/sql"
	"fmt"

	sqldriver "github.com/go-sql-driver/mysql"
)

type Config struct {
	Debug        bool
	DBType       string
	DSN          string
	MaxLifetime  int
	MaxOpenConns int
	MaxldleConns int
	TablePrefix  string
}

func CreateDatabaseWithMysql(cfg *sqldriver.Config) error {
	dsn := fmt.Sprintf("%s:%s@(%s)", cfg.User, cfg.Passwd, cfg.Addr)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	query := fmt.Sprintf("create database  if not exists `%s`   default  character set =`utf8mb4` ;", cfg.DBName)
	_, err = db.Exec(query)
	return err
}
