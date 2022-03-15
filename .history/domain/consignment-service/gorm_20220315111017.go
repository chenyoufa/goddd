package consignment_service

import (
	"demo1/domain/consignment-service/pkg/gormx"

	"gorm.io/gorm"
)

func InitGormDB() (*gorm.DB, func(), error) {
	db, err := NewGormDB()
	if err != nil {
		return nil, nil, err
	}

}

func NewGormDB() (*gorm.DB, error) {

	return gormx.New(&gormx.Config{
		Debug:        false,
		DBType:       "mysql",
		DSN:          "",
		MaxIdleConns: 1,
		MaxLifetime:  1,
		MaxOpenConns: 1,
		TablePrefix:  "",
	})

}
