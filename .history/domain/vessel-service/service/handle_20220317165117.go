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
	log.Println("vReq~~:", p[0])
	if err != nil {
		return err
	}
	if len(p) > 0 {
		temp := out.Vessel
		copier.Copy(temp, p[0])
		log.Println("vReqdd~~:", temp)
	}

	return nil
}
