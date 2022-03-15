package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	mysqlDriver "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
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

var golbalDB *gorm.DB

func (c *Config) Connect() (*gorm.DB, error) {
	var dialector gorm.Dialector

	switch strings.ToLower(c.DBType) {
	case "mysql":
		cfg, err := mysqlDriver.ParseDSN(c.DSN)
		if err != nil {
			log.Printf("ParseDSN fail err:%v", err)
			return nil, err
		}
		err = createDatabaseWithMySQL(cfg)
		if err != nil {
			log.Printf("create database width mysql fail err:%v", err)
			return nil, err
		}
		dialector = mysql.Open(c.DSN)
	case "postgres":
		dialector = postgres.Open(c.DSN)
	default:
		dialector = sqlite.Open(c.DSN)
	}
	gconfig := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   c.TablePrefix,
			SingularTable: true,
		},
	}
	db, err := gorm.Open(dialector, gconfig)
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
	// db.AutoMigrate()
	golbalDB = db
	return db, nil
}

func GetDB(ctx context.Context) *gorm.DB {
	return golbalDB.WithContext(ctx)
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

func init() {

}
