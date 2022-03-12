package main

import(
	pb "demo1/micro/consignment-service/proto/consignment"
)

type IRepository struct{
	Create(p *pb.Consignment)
}

type Repository struct{
	consignments []*pb.Consignment
}