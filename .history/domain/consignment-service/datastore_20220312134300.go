package main

import (
	"strings"

	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
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
	switch strings.ToLower(c.DBType) {

	case "mysql":
		mysql.Open()
	case "postgres":
		postgres.Open()
	default:
		sqlite.Open()
	}

}
