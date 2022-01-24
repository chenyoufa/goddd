package gormgo

import (
	"gorm.io/gorm"
)

func New(db *gorm.DB) *Hook {
	db.AutoMigrate(new(Logger))

}
