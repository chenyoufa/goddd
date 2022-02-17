package gormx

import (
	"database/sql"
	"fmt"

	mysqlDriver "github.com/go-sql-driver/mysql"
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

func createDatabaseWithMySQL(cfg *mysqlDriver.Config) error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/", cfg.u)
	db, err := sql.Open()
}
