package main

import (
	"context"
	pb "demo1/consignment-service/proto/consignment"
	"log"
	"net"

	"google.golang.org/grpc"
)

type IRepository interface {
	Create(consignment *pb.Consignment) (*pb.Consignment, error)
}

type Repository struct {
	consignments []*pb.Consignment
}

func (r *Repository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
	r.consignments = append(r.consignments, consignment)
	return consignment, nil
}

func (r *Repository) GetAll() []*pb.Consignment {
	return r.consignments
}

type Service struct {
	rep Repository
}

func (s *Service) CreateConsignment(ctx context.Context, in *pb.Consignment) (*pb.Response, error) {
	news, err := s.rep.Create(in)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &pb.Response{Created: true, Consignment: news}, nil
}

func (s *Service) GetConsignments(ctx context.Context, req *pb.GetRequest) (*pb.Response, error) {
	allconsignments := s.rep.GetAll()
	resp := &pb.Response{Consignments: allconsignments}
	return resp, nil
}

func main() {

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Println(err)
	}
	server := grpc.NewServer()
	repo := Repository{}
	pb.RegisterShippingServerServer(server, &Service{repo})
	log.Println("run serve  ...")

	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to serve:%v", err)
	}

}
