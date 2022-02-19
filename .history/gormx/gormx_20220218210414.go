package gormx

import (
	"database/sql"
	"fmt"

	sqldriver "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
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
func (cfg *Config) New() (*gorm.DB, error) {

	switch cfg.DBType {

	case "mysql":
		cfg, err := sqldriver.ParseDSN(cfg.DSN)
		if err != nil {
			return nil, err
		}
		err = CreateDatabaseWithMysql(cfg)
		if err != nil {
			return nil, err
		}
		dialector := mysql.Open("mysql", cfg.DSN)

	case "postgres":

	default:

	}

}
