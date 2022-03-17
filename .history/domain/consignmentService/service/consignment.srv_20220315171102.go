package service

import (
	"context"
	"demo1/domain/consignmentService/dao"
	"demo1/domain/consignmentService/dao/consignment"
)

type ConsignmentSrv struct {
	consignmentRepo *dao.ConsignmentPepo
}

func (a *ConsignmentSrv) Get(ctx context.Context, id uint64) (*consignment.Consignment, error) {
	a.consignmentRepo.Get(ctx, id)
}
