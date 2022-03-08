package main

import (
	pb "demo1/consignment-service/proto/consignment"
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
	consignment []*pb.Consignment
}

func (repo *Repository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
	repo.consignment = append(repo.consignment, consignment)
	return consignment, nil
}

func (repo *Repository) GetAll() []*pb.Consignment {
	return repo.consignments
}