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
	MaxLifeTime  time.Time
	MaxdeleTime  time.Time
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
		return nil
	}
}

func CreateDatabaseMysql(config *sqldriver.Config) {

}
