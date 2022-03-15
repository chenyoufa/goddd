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
	Description string       `gorm:"size:6;uniqueIndex;default:'';not null;"`
	Weight      int32        ``
	Containers  []*Container ``
	VesselId    string       ``
}
type Container struct {
	util.Model
	UserId      string ``
	CoumenterId string ``
	Origin      string ``
}
