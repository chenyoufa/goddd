package main

import (
	pb "demo1/micro/vessel-service/proto"
)

type IRepository interface {
	Find(spec *pb.Specification) *pb.Response
}
