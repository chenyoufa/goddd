package service

import (
	"context"
	"demo1/domain/consignmentService/dao"
)

type ConsignmentSrv struct {
	consignmentRepo *dao.ConsignmentPepo
}

func (a *ConsignmentSrv) Get(ctx context.Context, id uint64) (*Consignment, error) {

}
