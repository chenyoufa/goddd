package dao

import (
	"strings"

	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) error {
	var dbType = "mysql"
	if strings.ToLower(dbType) == "mysql" {
		db = db.Set("gorm:table_options", "ENGINE=innoDB")
	}
}
