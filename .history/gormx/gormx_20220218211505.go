package gormx

import (
	"database/sql"
	"fmt"

	sqldriver "github.com/go-sql-driver/mysql"
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
	MaxldleConns int
	TablePrefix  string
}

func CreateDatabaseWithMysql(cfg *sqldriver.Config) error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)", cfg.User, cfg.Passwd, cfg.Addr)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	query := fmt.Sprintf("create database  if not exists `%s`   default  character set =`utf8mb4` ;", cfg.DBName)
	_, err = db.Exec(query)
	return err
}
func (config *Config) New() (*gorm.DB, error) {
	var dialector = gorm.Dialector
	switch config.DBType {

	case "mysql":
		cfg, err := sqldriver.ParseDSN(config.DSN)
		if err != nil {
			return nil, err
		}
		err = CreateDatabaseWithMysql(cfg)
		if err != nil {
			return nil, err
		}
		dialector = mysql.Open(config.DSN)

	case "postgres":
		dialector = postgres.Open(config.DSN)
	default:
		dialector = sqlite.Open(config.DSN)
	}
	db, err := gorm.Open(dialector)
	if err != nil {
		return nil, err
	}
	if config.Debug {
		db = db.Debug()
	}
	sqldb, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqldb.SetMaxIdleConns()
	return db, nil
}
