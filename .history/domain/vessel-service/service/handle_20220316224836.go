package service

import (
	"context"
	"demo1/domain/vessel-service/dao"
	pb "demo1/domain/vessel-service/proto"
	vesselPb "demo1/micro/vessel-service/proto"

	"github.com/jinzhu/copier"
)

type VesselSrv struct {
	Repo         *dao.ConsignmentPepo
	VesselClient vesselPb.VesselServiceClient
}

func (srv *VesselSrv) FindAvailable(ctx context.Context, in *pb.Specification) (*pb.Response, error) {

}

func (srv *VesselSrv) GetAll(ctx context.Context, req *pb.GetRequest, resp *pb.Response) error {
	consignments, err := srv.Repo.GetAll()
	if err != nil {
		return err
	}
	copier.Copy(resp.Consignments, consignments)
	return nil
}
