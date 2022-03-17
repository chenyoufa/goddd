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

func (repo *VesselPepo) QueryByCapacityMaxWeight(Capacity int32, MaxWeight int32) ([]*Vessel, error) {
	var cs []*Vessel

	db := GetVesselDB(context.Background(), repo.DB)
	if Capacity > 0 {
		db = db.Where(" Capacity<= ? ", Capacity)
	}
	if MaxWeight > 0 {
		db = db.Where(" Max_Weight <=?", MaxWeight)
	}
	db.Find(&cs)
	return cs, errors.WithStack(db.Error)
}
