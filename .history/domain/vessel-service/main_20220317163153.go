package main

import (
	"demo1/domain/vessel-service/dao/vessel"
	"demo1/domain/vessel-service/gormx"
	"demo1/domain/vessel-service/service"
	"log"

	"github.com/micro/go-micro/v2"
)

func main() {

	// vessels := []*pb.Vessel{
	// 	{Id: "vessel001", Name: "Boaty McBoatface", MaxWeight: 200000, Capacity: 500},
	// }
	server := micro.NewService(
		micro.Name("go.micro.srv.vessel"),
		micro.Version("latest"),
	)

	server.Init()
	db, _, err := gormx.InitGormDB()
	if err != nil {
		log.Panic(err)
	}
	repo := &vessel.VesselPepo{
		DB: db,
	}
	srv := service.VesselSrv{Repo: repo}

	// p, err := srv.Repo.QueryByCapacityMaxWeight(3, 1300)
	// for _, item := range p {
	// 	log.Println(*item)
	// }

	// pb.RegisterVesselServiceHandler(server.Server(), &srv)

	// //启动service
	// if err := server.Run(); err != nil {
	// 	log.Panic("vessel-srv服务启动失败 ...")
	// }

}
