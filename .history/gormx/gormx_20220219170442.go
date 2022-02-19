package gormx

import (
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
}

func CreateDatabaseMysql(config *sqldriver.Config) {

}
