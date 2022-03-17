package service

import (
	"context"
	"demo1/domain/consignmentService/dao"
	"demo1/domain/consignmentService/dao/consignment"
	"errors"
	pb "demo1/domain/consignmentService/proto"
)

type ConsignmentSrv struct {
	Repo *dao.ConsignmentPepo
}


func(srv *ConsignmentSrv)CreateConsignment(ctx context.Context,req )



func (a *ConsignmentSrv) Get(ctx context.Context, id uint64) (*consignment.Consignment, error) {

	p, err := a.Repo.GetAll()
	if err != nil {
		return nil, err
	} else if p == nil {
		return nil, errors.New("data nil")
	}
	return p, nil
}
