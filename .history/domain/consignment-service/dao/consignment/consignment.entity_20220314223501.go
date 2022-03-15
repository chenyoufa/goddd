package consignment

import (
	"context"
	"demo1/domain/consignment-service/dao/util"

	"gorm.io/gorm"
)

func GetUserDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return util.GetDBWithModel(ctx, defDB, new(Consignment))
}

type Consignment struct {
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
