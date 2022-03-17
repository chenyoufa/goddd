package consignment

import (
	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type Repository interface {
	Create(*Consignment) error
	GetAll() ([]*Consignment, error)
}

type ConsignmentPepo struct {
	DB *gorm.DB
}

func (repo *ConsignmentPepo) Create(c *Consignment) error {

	result := GetConsignmentDB(context.Background(), repo.DB).Create(c)
	return errors.WithStack(result.Error)
}

func (repo *ConsignmentPepo) GetAll() ([]*Consignment, error) {
	var cs []*Consignment
	result := GetConsignmentDB(context.Background(), repo.DB).Find(&cs)
	return cs, errors.WithStack(result.Error)
}

func (repo *ConsignmentPepo) Get(wherestr string) ([]*Consignment, error) {
	var cs []*Consignment
	result := GetConsignmentDB(context.Background(), repo.DB).Find(&cs).Where(wherestr)
	return cs, errors.WithStack(result.Error)
}
