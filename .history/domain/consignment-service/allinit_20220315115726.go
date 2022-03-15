package consignment_service

func Build() {
	db, cleanup2, err := InitGormDB()
	if err != nil {
		return nil, nil, nil
	}
}
