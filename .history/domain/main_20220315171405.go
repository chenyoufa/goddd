package main

import (
	"context"
	a "demo1/domain/consignmentService"
	"demo1/domain/consignmentService/dao/consignment"
	"demo1/domain/consignmentService/service"
	"fmt"
)

func main() {
	fmt.Println("start main")
	db, _, err := a.InitGormDB()
	if err != nil {
		fmt.Println(err)
		// return nil, nil, nil
	}
	consignmentRepo := &consignment.ConsignmentPepo{
		DB: db,
	}
	consignmentSrv := &service.ConsignmentSrv{
		consignmentRepo:consignmentRepo
	}
	p, err := consignmentRepo.Get(context.Background(), 1)
	fmt.Println(p.Description)
}
