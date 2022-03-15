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
	Description string       `gorm:"size:20;uniqueIndex;default:'';not null;"`
	Weight      int32        `gorm:"index;default:0"`
	Containers  []*Container ``
	VesselId    string       ``
}
type Container struct {
	util.Model
	UserId      string ``
	CoumenterId string ``
	Origin      string ``
}
