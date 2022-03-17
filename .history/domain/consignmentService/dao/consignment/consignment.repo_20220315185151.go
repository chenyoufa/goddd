package consignment

import (
	"demo1/domain/consignmentService/dao/util"

	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type Repository interface {
	Create(*Consignment) error
	GetAll() ([]*Consignment, error)
	Close()
}

type ConsignmentPepo struct {
	DB *gorm.DB
}

func (repo *ConsignmentPepo) Create(c *Consignment)

func (c *ConsignmentPepo) Get(ctx context.Context, id uint64) (*Consignment, error) {
	var item Consignment
	ok, err := util.FindOne(ctx, GetConsignmentDB(ctx, c.DB).Where("id=?", id), &item)
	if err != nil {
		return nil, errors.WithStack(err)
	} else if !ok {
		return nil, nil
	}
	return &item, nil
}