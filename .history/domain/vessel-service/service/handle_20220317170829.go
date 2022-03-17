package service

import (
	"context"
	"log"

	"demo1/domain/vessel-service/dao"
	pb "demo1/domain/vessel-service/proto"

	"github.com/jinzhu/copier"
)

type VesselSrv struct {
	Repo *dao.VesselPepo
}

func (srv *VesselSrv) FindAvailable(ctx context.Context, in *pb.Specification, out *pb.Response) error {

	p, err := srv.Repo.QueryByCapacityMaxWeight(in.Capacity, in.MaxWeight)
	if err != nil {
		return err
	}
	if len(p) > 0 {
		// temp := &pb.Vessel{}
		temp := out.Vessel
		copier.Copy(temp, p[0])
		// out.Vessel = temp
	}
	log.Println("vReqdd~~:", out.Vessel)
	return nil
}
