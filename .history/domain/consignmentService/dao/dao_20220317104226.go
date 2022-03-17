package dao

import (
	"demo1/domain/consignmentService/dao/consignment"
	"strings"

	"gorm.io/gorm"
)

type (
	ConsignmentPepo = consignment.Repository
)

func AutoMigrate(db *gorm.DB) error {
	var dbType = "mysql"
	if strings.ToLower(dbType) == "mysql" {
		db = db.Set("gorm:table_options", "ENGINE=innoDB")
	}
	return db.AutoMigrate(
		new(consignment.Consignment),
	)
}
