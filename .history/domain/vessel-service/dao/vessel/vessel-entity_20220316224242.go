package vessel

import (
	"context"
	"demo1/domain/vessel-service/dao/util"

	"gorm.io/gorm"
)

func GetVesselDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return util.GetDBWithModel(ctx, defDB, new(Vessel))
}

type Vessel struct {
	util.Model
	Capacity  int32  `gorm:"not null "`
	MaxWeight int32  `gorm:"not null"`
	Name      string `gorm:"size:20;not null"`
	Available bool   `gorm:"default:false"`
	OwerD     string `gorm:"size:20;not null"`
}
type Container struct {
	util.Model
	UserId        string `gorm:"size:20;not null;"`
	CoumenterId   string `gorm:"size:20;not null"`
	Origin        string `gorm:"size:20;not null"`
	ConsignmentID uint
}
