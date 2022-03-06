package main

import (
	"context"
	pb "demo1/consignment-service/proto/consignment"
	"log"
	"net"

	"google.golang.org/grpc"
)

const (
	PORT = ":50051"
)

//仓库接口
type IRespository interface {
	Create(consignment *pb.Consignment) (*pb.Consignment, error) //存放新货物
}

//我们存放多批货物的仓库，实现了IRepository接口
type Repository struct {
	consignments []*pb.Consignment
}

func (repo *Repository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
	repo.consignments = append(repo.consignments, consignment)
	return consignment, nil
}

func (repo *Repository) GetAll() []*pb.Consignment {
	return repo.consignments
}

//定义微服务
type service struct {
	repo Repository
}

func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment) (*pb.Response, error) {
	consignment, err := s.repo.Create(req)
	if err != nil {
		return nil, err
	}
	resp := &pb.Response{Created: true, Consignment: consignment}
	return resp, nil
}

func main() {
	listener, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatalf("failed to listen:%v", err)
	}
	log.Printf("listen on:%s\n", PORT)
	server := grpc.NewServer()
	repo := Repository{}
	pb.RegisterShippingServiceServer(server, &service{repo})
	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to serve:%v", err)
	}
}
