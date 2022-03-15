package consignment_service

func main() {
	db, cleanup2, err := InitGormDB()
	if err != nil {
		// return nil, nil, nil
	}
	consignmentRepo := &consignment.ConsignmentPepo{
		DB: db,
	}
}
