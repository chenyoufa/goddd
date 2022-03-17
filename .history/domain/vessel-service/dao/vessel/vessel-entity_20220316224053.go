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
	Capacity  int32  `gorm:"index;not null "`
	MaxWeight int32  ``
	Name      string ``
	Available bool   ``
	OwerD     string ``
}
type Container struct {
	util.Model
	UserId        string `gorm:"size:20;not null;"`
	CoumenterId   string `gorm:"size:20;not null"`
	Origin        string `gorm:"size:20;not null"`
	ConsignmentID uint
}
