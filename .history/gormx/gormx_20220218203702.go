package gormx

import (
	"database/sql"

	sqldriver "github.com/go-sql-dirver/mysql"
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

func CreateSqlWithMysql(config *sqldriver.Config) {
	sql.Open()
	sqldriver
}
