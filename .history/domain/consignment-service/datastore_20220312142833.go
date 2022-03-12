package main

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	mysqlDriver "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Config struct {
	Debug        bool
	DBType       string
	DSN          string
	MaxLifetime  int
	MaxOpenConns int
	MaxIdleConns int
	TablePrefix  string
}

func (c *Config) New() (*gorm.DB, error) {
	var dialector gorm.Dialector
	cfg, err := mysqlDriver.ParseDSN(c.DSN)

	switch strings.ToLower(c.DBType) {
	case "mysql":

		dialector = mysql.Open(c.DSN)
	case "postgres":
		dialector = postgres.Open(c.DSN)
	default:
		dialector = sqlite.Open(c.DSN)
	}

	db, err := gorm.Open(dialector)
	if err != nil {
		return nil, err
	}
	if c.Debug {
		db = db.Debug()
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxOpenConns(c.MaxIdleConns)
	sqlDB.SetMaxIdleConns(c.MaxIdleConns)
	sqlDB.SetConnMaxLifetime(time.Duration(c.MaxLifetime))
	return db, nil
}

func createDatabaseWithMySQL(cfg *mysqlDriver.Config) error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/", cfg.User, cfg.Passwd, cfg.Addr)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	defer db.Close()
	query := fmt.Sprintf("create database if not exits `%s` default  character set =`utf8mb4`;", cfg.DBName)
	_, err = db.Exec(query)
	return err
}
