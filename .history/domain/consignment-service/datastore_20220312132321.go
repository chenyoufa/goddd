package main

import (
	mysqlDriver "github.com/go-sql-driver/mysql"
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

func CreateSession(host string) {
	var dialector gorm.Dialector
	cfg, err := mysqlDriver.ParseDSN()

}
