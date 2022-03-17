package service

import (
	"context"
	"demo1/domain/consignmentService/dao"
	"demo1/domain/consignmentService/dao/consignment"
	"errors"
)

type ConsignmentSrv struct {
	consignmentRepo *dao.ConsignmentPepo
}

func (a *ConsignmentSrv) Get(ctx context.Context, id uint64) (*consignment.Consignment, error) {
	p, err := a.consignmentRepo.Get(ctx, id)
	if err != nil {
		return nil, err
	} else if p == nil {
		return nil, errors.ErrNotFound
	}
}
