package consignment_service

import (
	"demo1/domain/consignment-service/dao/consignment"
)

func main() {
	db, _, err := InitGormDB()
	if err != nil {
		// return nil, nil, nil
	}
	consignmentRepo := &consignment.ConsignmentPepo{
		DB: db,
	}
}
