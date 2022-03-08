package main

import (
	"context"
	pb "demo1/one1/consignment-service/proto/consignment"
	"log"
	"net"

	"google.golang.org/grpc"
)

type IRepository interface {
	Create(consignment *pb.Consignment) (*pb.Consignment, error)
	GetAll() ([]*pb.Consignment, error)
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

func (s *Service) CreateConsignment(ctx context.Context, consignment *pb.Consignment) (*pb.Response, error) {
	newconsignment, err := s.rep.Create(consignment)
	if err != nil {
		log.Printf("create consignment fail err:%v", err)
		return nil, err
	}
	return &pb.Response{Created: true, Consignment: newconsignment}, nil

}
func (s *Service) GetAll(ctx context.Context, req *pb.GetRequest) (*pb.Response, error) {
	consignments := s.rep.GetAll()
	return &pb.Response{Consignments: consignments}, nil
}

func main() {

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Printf("listen port fail err:%v", err)
	}
	log.Println("server run ... \n")
	server := grpc.NewServer()

	rep := Repository{}

	pb.RegisterShippingServiceServer(server, &Service{rep})

	if err := server.Serve(listener); err != nil {
		log.Printf("server fail err:%v", err)
	}

}
