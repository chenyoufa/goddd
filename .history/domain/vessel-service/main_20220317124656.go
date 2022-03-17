package main

import (
	"demo1/domain/consignmentService/service"
	pb "demo1/domain/vessel-service/proto"

	"github.com/micro/go-micro/v2"
)

func main() {

	vessels := []*pb.Vessel{
		{Id: "vessel001", Name: "Boaty McBoatface", MaxWeight: 200000, Capacity: 500},
	}
	server := micro.NewService(
		micro.Name("go.micro.srv.vessel"),
		micro.Version("latest"),
	)
	server.Init()
	pb.RegisterVesselServiceHandler(server.Server(), &service.VesselSrv{})

}
