package main

import (
	"context"
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
	return nil, errors.New("not find")
}

type Service struct {
	rep Repository
}

func (s *Service) FindAvailable(ctx context.Context, spec *pb.Specification, res *pb.Response) error {
	p, err := s.rep.FindAvailable(spec)
	if err != nil {
		return err
	}
	res.Vessel = p
	return nil
}
