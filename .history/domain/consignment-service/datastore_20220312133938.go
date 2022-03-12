package main

import (
	"strings"

	"github.com/jinzhu/gorm"
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

	switch strings.ToLower(c.DBType) {

	case "mysql":
		mysql
	case "sqlserver":

	default:

	}

}
