package main

import (
	a "demo1/domain/consignmentService"
	"demo1/domain/consignmentService/dao/consignment"
	"demo1/domain/consignmentService/service"
	"fmt"

	micro "github.com/micro/go-micro/v2"
)

const (
	DEFAULT_HOST = "localhost:27017"
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

}
