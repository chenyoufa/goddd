package service

import (
	"context"

	"demo1/domain/vessel-service/dao"
	pb "demo1/domain/vessel-service/proto"
	vesselPb "demo1/micro/vessel-service/proto"

	"github.com/jinzhu/copier"
)

type VesselSrv struct {
	Repo         *dao.VesselPepo
	VesselClient vesselPb.VesselServiceClient
}

func (srv *VesselSrv) FindAvailable(ctx context.Context, in *pb.Specification, out *pb.Response) error {

	p, err := srv.Repo.QueryByCapacityMaxWeight(in.Capacity, in.MaxWeight)
	if err != nil {
		return err
	}
	copier.Copy(out.Vessel, p)
	return nil
}
