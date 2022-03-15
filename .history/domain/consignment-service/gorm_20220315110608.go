package consignment_service

import (
	"demo1/domain/consignment-service/pkg/gormx"

	"gorm.io/gorm"
)

func InitGormDB() (*gorm.DB, func(), error) {

}

func NewGormDB() (*gorm.DB, error) {

	return gormx.New(&gormx.Config{
		Debug: false,
	})

}
