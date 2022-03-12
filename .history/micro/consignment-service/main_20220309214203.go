package main

import (
	pb "demo1/micro/consignment-service/proto/consignment"
	"log"
)

type IRepository interface {
	Create(consignment *pb.Consignment) (*pb.Consignment, error)
}

type Repository struct {
	consignments []*pb.Consignment
}

func (rep *Repository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
	rep.consignments = append(rep.consignments, consignment)
	return consignment, nil
}

func (rep *Repository) GetAll() []*pb.Consignment {
	return rep.consignments
}

type Service struct {
	rep Repository
}

func (s *Service) CreateConsignment(new_consignment *pb.Consignment) (*pb.Response, error) {
	p, err := s.rep.Create(new_consignment)
	if err != nil {
		log.Printf("create consignment fail err:%v", err)
		return nil, err
	}
	return &pb.Response{Created: true, Consignment: p}, nil

}
