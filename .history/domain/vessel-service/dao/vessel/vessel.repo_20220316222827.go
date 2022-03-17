package vessel

import (
	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type Repository interface {
	Create(*Consignment) error
	GetAll() ([]*Consignment, error)
}

type VesselPepo struct {
	DB *gorm.DB
}

func (repo *VesselPepo) Create(c *Consignment) error {

	result := GetConsignmentDB(context.Background(), repo.DB).Create(c)
	return errors.WithStack(result.Error)
}

func (repo *VesselPepo) GetAll() ([]*Consignment, error) {
	var cs []*Consignment
	result := GetConsignmentDB(context.Background(), repo.DB).Find(&cs)
	return cs, errors.WithStack(result.Error)
}
