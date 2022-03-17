package service

import (
	"context"
	"demo1/domain/consignmentService/dao"
	"demo1/domain/consignmentService/dao/consignment"
	pb "demo1/domain/consignmentService/proto"
	"log"
	"time"
)

type ConsignmentSrv struct {
	Repo *dao.ConsignmentPepo
}

func (srv *ConsignmentSrv) CreateConsignment(ctx context.Context, req *pb.Consignment, resp *pb.Response) error {
	var a *consignment.Consignment
	a.CreatedAt = time.Now()
	a.Description = req.Description
	a.Weight = req.Weight
	// for _,item:=range req.Containers{
	// 	a.Containers=append(a.Containers, *item.)
	// }

	err := srv.Repo.Create(a)
	if err != nil {
		log.Println(err)
		return err
	}
	resp.Created = true
	return nil
}
