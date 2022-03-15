package consignment_service

import "gorm.io/gorm"

func InitGormDB() (*gorm.DB, func(), error) {

}

func NewGormDB() (*gorm.DB, error) {

	return gormx.New()

}
