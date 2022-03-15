package consignment_service

import (
	"demo1/domain/consignment-service/dao/consignment"
)

func main() {
	db, cleanup2, err := InitGormDB()
	if err != nil {
		// return nil, nil, nil
	}
	consignmentRepo := &consignment.ConsignmentPepo{
		DB: db,
	}
}
