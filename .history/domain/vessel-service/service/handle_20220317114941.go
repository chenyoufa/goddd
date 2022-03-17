package service

import (
	"context"

	"demo1/domain/vessel-service/dao"
	pb "demo1/domain/vessel-service/proto"
	vesselPb "demo1/micro/vessel-service/proto"
)

type VesselSrv struct {
	Repo         *dao.VesselPepo
	VesselClient vesselPb.VesselServiceClient
}

func (srv *VesselSrv) FindAvailable(ctx context.Context, in *pb.Specification, out *pb.Response) error {
	wherestr := ""
	p, err := srv.Repo.Get(wherestr)
	if err != nil {
		return err
	}
	res.Vessel = p.Find()
	return nil
}
