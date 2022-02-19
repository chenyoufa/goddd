package gormx

import (
	"strings"
	"time"

	sqldriver "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	Debug        bool
	DbType       string
	DSN string
	MaxOpenConns int
	MaxLifeConns int
	MaxLifeTime  time.Time
	MaxdeleTime  time.Time
}

func (c *Config) New() *gorm.DB {

	var dialector gorm.Dialector
	switch strings.ToLower(c.DbType) {
	case "mysql":
		dialector=mysql.Open(c.)
	case "postgres":

	default:

	}

}

func CreateDatabaseMysql(config *sqldriver.Config) {

}
