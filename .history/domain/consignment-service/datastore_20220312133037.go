package main

import (
	mysqlDriver "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
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

func New(c *Config)New()(*gorm){


}
