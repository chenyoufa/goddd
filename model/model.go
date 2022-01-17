package model

import (
	"time"

	"gorm.io/gorm"
)

type BASE_MODEL struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
