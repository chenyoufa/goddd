package consignment

import (
	"context"
	"demo1/domain/consignmentService/dao/util"

	"gorm.io/gorm"
)

func GetUserDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return util.GetDBWithModel(ctx, defDB, new(Consignment))
}

type Consignment struct {
	util.Model
	Description string `gorm:"size:20;not null;"`
	Weight      int32  `gorm:"index;not null"`
	Containers  []Container
	VesselId    string `gorm:"size:20;not null"`
}
type Container struct {
	util.Model
	UserId      string `gorm:"size:20;not null;"`
	CoumenterId string `gorm:"size:20;not null"`
	Origin      string `gorm:"size:20;not null"`
}
