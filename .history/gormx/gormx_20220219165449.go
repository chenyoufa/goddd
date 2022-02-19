package gormx

import (
	"time"

	"gorm.io/gorm"
)

type Config struct {
	Debug        bool
	DbType       string
	MaxOpenConns int
	MaxLifeConns int
	MaxLifeTime  time.Duration
	MaxdeleTime  time.Time
}

func New() *gorm.DB {

}
