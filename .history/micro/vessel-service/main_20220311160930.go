package main

import (
	pb "demo1/micro/vessel-service/proto"
	"errors"
)

type IRepository interface {
	FindAvailable(spec *pb.Specification) (*pb.Vessel, error)
}
type Repository struct {
	vessels []*pb.Vessel
}

func (rep *Repository) FindAvailable(spec *pb.Specification) (*pb.Vessel, error) {
	for _, item := range rep.vessels {
		if spec.Capacity <= item.Capacity && spec.MaxWeight <= item.MaxWeight {
			return item, nil
		}
	}
	return nil, errors.New("not file ")
}
