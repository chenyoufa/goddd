package service

import (
	"context"
	"demo1/domain/consignmentService/dao"
	pb "demo1/domain/consignmentService/proto"
)

type ConsignmentSrv struct {
	Repo *dao.ConsignmentPepo
}

func (srv *ConsignmentSrv) CreateConsignment(ctx context.Context, req *pb.Consignment, resp *pb.Response) error {

}
