package main

import (
	"strings"

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

}
