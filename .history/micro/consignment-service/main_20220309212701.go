package main

import (
	pb "demo1/micro/consignment-service/proto/consignment"
)

type IRepository interface {
	Create(consignment *pb.Consignment) *pb.Consignment
}

type Repository struct {
	consignments []*pb.Consignment
}
