package main

import (
	"context"
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

type Service struct {
	rep Repository
}

func (s *Service) CreateConsignment(ctx context.Context, newp *pb.Consignment) (*pb.Response, error) {
	s.rep.Create(newp)
	return &pb.Response{Created: true, Consignments: s.rep.consignments}, nil
}
func (s *Service) GetConsignments(context.Context, *pb.Request) (*pb.Response, error) {

}

func main() {

}
