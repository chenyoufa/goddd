package main

import (
	pb "demo1/consignment-service/proto/consignment"
)

const (
	PORT = "8080"
)

type IRepository interface {
	Create(p *pb.Consignment) *pb.Consignment
}

type Repository struct {
	consignments []*pb.Consignment
}

func (rep *Repository) Create(p *pb.Consignment) *pb.Consignment {
	rep.consignments = append(rep.consignments, p)
	return p
}
func (rep *Repository) GetAll() []*pb.Consignment {
	return rep.consignments
}

func main() {

}
