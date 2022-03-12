package main

import (
	pb "demo1/micro/vessel-service/proto"
)

type IRepository interface {
	FindSpecification(spec *pb.Specification) (*pb.Vessel,error)
}
type Repository struct{
	vessels []*pb.Vessel
}

func
