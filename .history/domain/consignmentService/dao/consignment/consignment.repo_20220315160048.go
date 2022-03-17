package consignment

import (
	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type ConsignmentPepo struct {
	DB *gorm.DB
}

func (c *ConsignmentPepo) Get(ctx context.Context) {
	var item User
	ok, err := util.FindOne(ctx, GetConsignmentDB(ctx, a.DB).Where("id=?", id), &item)
	if err != nil {
		return nil, errors.WithStack(err)
	} else if !ok {
		return nil, nil
	}
}
