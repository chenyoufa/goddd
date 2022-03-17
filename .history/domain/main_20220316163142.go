package main

import (
	a "demo1/domain/consignmentService"
	"demo1/domain/consignmentService/dao/consignment"
	"demo1/domain/consignmentService/service"
	"fmt"
	"log"

	pb "demo1/domain/consignmentService/proto"

	micro "github.com/micro/go-micro/v2"
)

func main() {
	fmt.Println("start main")
	server := micro.NewService(
		micro.Name("go.micro.srv.consignment"),
		micro.Version("latest"),
	)
	server.Init()
	db, _, err := a.InitGormDB()
	if err != nil {
		fmt.Println(err)
		// return nil, nil, nil
	}
	consiRepo := &consignment.ConsignmentPepo{
		DB: db,
	}
	consignmentSrv := &service.ConsignmentSrv{
		Repo: consiRepo,
	}

	pb.RegisterShippingServiceHandler(server.Server(), consignmentSrv)

	if err := server.Run(); err != nil {
		log.Fatalf("failed to serve:%v", err)
	}
}
