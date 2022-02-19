package gormx

import (
	"fmt"
	"strings"
	"time"

	sqldriver "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Config struct {
	Debug        bool
	DbType       string
	DSN          string
	MaxOpenConns int
	MaxLifeConns int
	MaxLifeTime  time.Duration
	MaxIdleTime  time.Duration
	TablePrefix  string
}

func (c *Config) New() (*gorm.DB, error) {

	var dialector gorm.Dialector
	switch strings.ToLower(c.DbType) {
	case "mysql":

		conf, err := sqldriver.ParseDSN(c.DSN)
		if err != nil {
			return nil, err
		}
		err := CreateDatabaseMysql(conf)
		if err != nil {
			return err
		}

		dialector = mysql.Open(c.DSN)
	case "postgres":
		dialector = postgres.Open(c.DSN)
	default:
		dialector = sqlite.Open(c.DSN)
	}
	gormconfig := &gorm.Config{NamingStrategy: schema.NamingStrategy{
		TablePrefix:   c.TablePrefix,
		SingularTable: true,
	}}
	db, err := gorm.Open(dialector, gormconfig)
	if err != nil {
		return nil, err
	}
	sqldb, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqldb.SetConnMaxLifetime(c.MaxLifeTime)
	sqldb.SetConnMaxIdleTime(c.MaxIdleTime)
	sqldb.SetMaxOpenConns(c.MaxOpenConns)

	return db, nil
}

func CreateDatabaseMysql(config *sqldriver.Config) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)", config.User, config.Passwd, config.Addr)
	dialector := mysql.Open(dsn)
	db, err := gorm.Open(dialector)
	if err != nil {
		return err
	}
	query := fmt.Sprintf("create database if not existis `%s` default character set=`utf8mb4`;")
}
