package main

import (
	"demo1/domain/consignmentService/dao"
	"demo1/domain/consignmentService/pkg/gormx"

	"gorm.io/gorm"
)

func DDD() int {
	return 1
}
func LnitGormDB() (*gorm.DB, func(), error) {
	db, err := NewGormDB()
	if err != nil {
		return nil, nil, err
	}
	cleanFunc := func() {}
	EnableAutoMigrate := true
	if EnableAutoMigrate {
		err = dao.AutoMigrate(db)
		if err != nil {
			return nil, cleanFunc, err
		}
	}
	return db, cleanFunc, nil
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
