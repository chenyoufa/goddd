package main

import (
	mysqlDriver "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

func CreateSession(host string) {
	var dialector gorm.Dialector
	cfg, err := mysqlDriver.ParseDSN()

}
