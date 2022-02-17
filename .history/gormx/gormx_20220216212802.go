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
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/", cfg.User, cfg.Passwd, cfg.Addr)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	defer db.Close()
	query := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s` DEFAULT character set=`utf8mb4`;", cfg.DBName)
}
