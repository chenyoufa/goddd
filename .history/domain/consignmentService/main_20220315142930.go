package main

import (
	"context"
	"demo1/domain/consignmentService/dao/consignment"
	"fmt"
)

func main() {
	db, err := NewGormDB()
	if err != nil {

	}

	db, _, err := LnitGormDB()
	if err != nil {
		fmt.Println(err)
		// return nil, nil, nil
	}
	consignmentRepo := &consignment.ConsignmentPepo{
		DB: db,
	}
	consignmentRepo.Get(context.Background())

}
