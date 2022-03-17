package dao

import (
	"demo1/domain/vessel-service/dao/vessel"
	"strings"

	"gorm.io/gorm"
)

type (
	VesselPepo = vessel.VesselPepo
)

func AutoMigrate(db *gorm.DB) error {
	var dbType = "mysql"
	if strings.ToLower(dbType) == "mysql" {
		db = db.Set("gorm:table_options", "ENGINE=innoDB")
	}
	return db.AutoMigrate(
		new(vessel.Vessel),
	)
}
