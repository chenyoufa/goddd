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
	Id          string       ` `
	Description string       ``
	Weight      int32        ``
	Containers  []*Container ``
	VesselId    string       ``
}
type Container struct {
	Id          string ``
	UserId      string ``
	CoumenterId string ``
	Origin      string ``
}
