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

	vReq := &vesselPb.Specification{
		Capacity:  int32(len(req.Containers)),
		MaxWeight: req.Weight,
	}
	// log.Println("vReq~~:", vReq)
	_, err := srv.VesselClient.FindAvailable(context.Background(), vReq)
	if err != nil {
		// log.Panic(err)
		return err
	}
	// log.Println("vResp~~:", vResp)
	req.VesselId = "10059" // vResp.Vessel.Id
	req.Containers = nil
	data := &consignment.Consignment{}
	pbq := &pb.Consignment{Id: "100", Description: "des", Weight: 10, VesselId: "10086"}

	log.Println("data:", pbq)
	log.Println("req:", req)

	copier.Copy(data, pbq)

	// for _,item:=range req.Containers{
	// 	a.Containers=append(a.Containers, *item.)
	// }
	err = srv.Repo.Create(data)
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
