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
	Id          string       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Description string       `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Weight      int32        `protobuf:"varint,3,opt,name=weight,proto3" json:"weight,omitempty"`
	Containers  []*Container `protobuf:"bytes,4,rep,name=containers,proto3" json:"containers,omitempty"`
	VesselId    string       `protobuf:"bytes,5,opt,name=vessel_id,json=vesselId,proto3" json:"vessel_id,omitempty"`
}
type Container struct {
	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId      string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	CoumenterId string `protobuf:"bytes,3,opt,name=coumenter_id,json=coumenterId,proto3" json:"coumenter_id,omitempty"`
	Origin      string `protobuf:"bytes,4,opt,name=origin,proto3" json:"origin,omitempty"`
}
