package main

import (
	"context"
	a "demo1/domain/consignmentService"
	"demo1/domain/consignmentService/dao/consignment"
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
	consignmentRepo.Get(context.Background())

}
