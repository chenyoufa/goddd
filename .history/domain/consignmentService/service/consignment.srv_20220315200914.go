package service

import (
	"context"
	"demo1/domain/consignmentService/dao"
	"demo1/domain/consignmentService/dao/consignment"
	"errors"
)

type ConsignmentSrv struct {
	Repo         *dao.ConsignmentPepo
	vesselClient vesselPb.VesselServiceClient
}

func (a *ConsignmentSrv) Get(ctx context.Context, id uint64) (*consignment.Consignment, error) {
	p, err := a.Repo.Get(ctx, id)
	if err != nil {
		return nil, err
	} else if p == nil {
		return nil, errors.New("data nil")
	}
	return p, nil
}
