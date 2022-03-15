package consignment_service

import (
	"context"
	"demo1/domain/consignment-service/dao/consignment"
	"fmt"
)

func main() {
	db, _, err := InitGormDB()
	if err != nil {
		fmt.Println(err)
		// return nil, nil, nil
	}
	consignmentRepo := &consignment.ConsignmentPepo{
		DB: db,
	}
	consignmentRepo.Get(context.Background())
}
