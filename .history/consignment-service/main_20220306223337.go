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
