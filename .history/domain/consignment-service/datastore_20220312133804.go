package main

import (
	"strings"

	"github.com/jinzhu/gorm"
	"gorm.io/driver/mysql"
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

	mysql.Open()

	switch strings.ToLower(c.DBType) {

	case "mysql":

	case "sqlserver":

	default:

	}

}
