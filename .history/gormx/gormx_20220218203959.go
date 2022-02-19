package gormx

import (
	"database/sql"
	"fmt"

	sqldriver "github.com/go-sql-driver/mysql"
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

func CreateSqlWithMysql(cfg *sqldriver.Config) error {
	dsn := fmt.Sprintf("%s:%s@(%s)", cfg.User, cfg.Passwd, cfg.Addr)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	query := ""
	return db.Exec(query).Erro
}
