package main

import (
	pb "demo1/micro/consignment-service/proto/consignment"
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
