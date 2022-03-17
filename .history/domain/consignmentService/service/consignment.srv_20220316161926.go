package service

import (
	"context"
	"demo1/domain/consignmentService/dao"
	"demo1/domain/consignmentService/dao/consignment"
	pb "demo1/domain/consignmentService/proto"
	"log"

	"github.com/jinzhu/copier"
)

type ConsignmentSrv struct {
	Repo *dao.ConsignmentPepo
}

func (srv *ConsignmentSrv) CreateConsignment(ctx context.Context, req *pb.Consignment, resp *pb.Response) error {
	var a *consignment.Consignment
	copier.Copy(a, req)
	// for _,item:=range req.Containers{
	// 	a.Containers=append(a.Containers, *item.)
	// }
	err := srv.Repo.Create(a)
	if err != nil {
		log.Println(err)
		return err
	}
	resp.Created = true
	resp.Consignment = req
	return nil
}

func (srv *ConsignmentSrv) GetConsignments(ctx context.Context, req *pb.GetRequest, resp *pb.Response) error {
	consignments, err := srv.Repo.GetAll()
	if err != nil {
		return err
	}
	copier.Copy(resp.Consignments, consignments)
	resp.Consignments = consignments
	return nil
}
