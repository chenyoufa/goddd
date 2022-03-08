package main

import (
	"context"
	pb "demo1/consignment-service/proto/consignment"
	"log"
	"net"

	"google.golang.org/grpc"
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
	return &pb.Response{Created: true}, nil
}
func (s *Service) GetConsignments(context.Context, *pb.Request) (*pb.Response, error) {
	return &pb.Response{Consignments: s.rep.consignments}, nil
}

func main() {

	listener, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Printf("listen fail  ", err)
	}
	log.Println("run port ", PORT)
	server := grpc.NewServer()

	pb.RegisterShippingServiceServer(server, &Service{rep: &Repository})
	if err := server.Serve(listener); err != nil {
		log.Printf("server run fail ", err)
	}

}
