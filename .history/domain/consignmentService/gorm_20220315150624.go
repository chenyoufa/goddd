package consignmentService

import (
	"demo1/domain/consignmentService/dao"
	"demo1/domain/consignmentService/pkg/gormx"
	"fmt"

	"gorm.io/gorm"
)

func LnitGormDB() (*gorm.DB, func(), error) {
	db, err := NewGormDB()
	if err != nil {
		fmt.Println(err)
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
		DSN:          "fage:Fage501526~@tcp(127.0.0.1)/mytest?charset=utf8&parseTime=True&loc=Local",
		MaxIdleConns: 1,
		MaxLifetime:  1,
		MaxOpenConns: 1,
		TablePrefix:  "",
	})

}
