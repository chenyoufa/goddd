package service

import (
	"context"
	"demo1/domain/consignmentService/dao"
	"demo1/domain/consignmentService/dao/consignment"
	pb "demo1/domain/consignmentService/proto"
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
	err := srv.Repo.Create(a)
}
