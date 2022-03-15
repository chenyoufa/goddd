package consignment

import (
	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type ConsignmentPepo struct {
	DB *gorm.DB
}

func (c *ConsignmentPepo) Get(ctx context.Context) {

}
