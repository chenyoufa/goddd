package vessel

import (
	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type Repository interface {
	Create(*Vessel) error
	GetAll() ([]*Vessel, error)
}

type VesselPepo struct {
	DB *gorm.DB
}

func (repo *VesselPepo) Create(c *Vessel) error {
	result := GetVesselDB(context.Background(), repo.DB).Create(c)
	return errors.WithStack(result.Error)
}

func (repo *VesselPepo) GetAll() ([]*Vessel, error) {
	var cs []*Vessel
	result := GetVesselDB(context.Background(), repo.DB).Find(&cs)
	return cs, errors.WithStack(result.Error)
}
func (repo *VesselPepo) Get(wherestr string) ([]*Vessel, error) {
	var cs []*Vessel
	result := GetVesselDB(context.Background(), repo.DB)
	if len(wherestr) > 0 {
		result = result.Where(wherestr)
	}
	result.Find(&cs)
	return cs, errors.WithStack(result.Error)
}
