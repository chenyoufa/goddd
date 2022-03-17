package main

import (
	"context"
	pb "demo1/micro/consignment-service/proto/consignment"
	"log"

	micro "github.com/micro/go-micro/v2"
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

func (s *Service) CreateConsignment(ctx context.Context, new_consignment *pb.Consignment, res *pb.Response) error {
	p, err := s.rep.Create(new_consignment)
	if err != nil {
		log.Printf("create consignment fail err:%v", err)
		return nil, err
	}
	return &pb.Response{Created: true, Consignment: p}, nil
}

func (s *Service) GetAll(context.Context, pb.GetRequest, *pb.Response) error {
	all := s.rep.GetAll()
	return &pb.Response{Consignments: all}, nil
}

func main() {

	server := micro.NewService(
		micro.Name("go.micro.srv.consignment"),
		micro.Version("latest"),
	)
	server.Init()
	repo := Repository{}
	pb.RegisterShippingServiceHandler(server.Server(), &Service{repo})
	if err := server.Run(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}