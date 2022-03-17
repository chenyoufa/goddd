package service

import (
	"context"
	"demo1/domain/consignment-service/dao"
	"demo1/domain/consignment-service/dao/consignment"
	pb "demo1/domain/consignment-service/proto"
	vesselPb "demo1/micro/vessel-service/proto"
	"log"

	"github.com/jinzhu/copier"
)

type ConsignmentSrv struct {
	Repo         *dao.ConsignmentPepo
	VesselClient vesselPb.VesselServiceClient
}

func (srv *ConsignmentSrv) CreateConsignment(ctx context.Context, req *pb.Consignment, resp *pb.Response) error {

	log.Println("weight~~:", req.Weight)
	vReq := &vesselPb.Specification{
		Capacity:  int32(len(req.Containers)),
		MaxWeight: req.Weight,
	}
	vResp, err := srv.VesselClient.FindAvailable(context.Background(), vReq)
	if err != nil {
		// log.Panic(err)
		return err
	}
	req.VesselId = vResp.Vessel.Id
	var a *consignment.Consignment
	copier.Copy(a, req)
	// for _,item:=range req.Containers{
	// 	a.Containers=append(a.Containers, *item.)
	// }
	err = srv.Repo.Create(a)
	if err != nil {
		// log.Panic(err)
		return err
	}
	resp.Created = true
	resp.Consignment = req
	return nil
}

func (srv *ConsignmentSrv) GetAll(ctx context.Context, req *pb.GetRequest, resp *pb.Response) error {
	consignments, err := srv.Repo.GetAll()
	if err != nil {
		return err
	}
	copier.Copy(resp.Consignments, consignments)
	return nil
}
