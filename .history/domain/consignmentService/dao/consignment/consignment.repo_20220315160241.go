package consignment

import (
	"demo1/domain/consignmentService/dao/util"

	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type ConsignmentPepo struct {
	DB *gorm.DB
}

func (c *ConsignmentPepo) Get(ctx context.Context) (bool, error) {
	var item Consignment
	ok, err := util.FindOne(ctx, GetConsignmentDB(ctx, a.DB).Where("id=?", id), &item)
	if err != nil {
		return nil, errors.WithStack(err)
	} else if !ok {
		return nil, nil
	}
}
