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

	if strings.ToLower(c.DBType) == "mysql" {

	} else if strings.ToLower(c.DBType) == "postgret" {

	}

}
