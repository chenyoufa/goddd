package main

import (
	a "demo1/domain/consignmentService"
	"demo1/domain/consignmentService/dao/consignment"
	pb "demo1/domain/consignmentService/proto"
	"demo1/domain/consignmentService/service"
	vesselPb "demo1/micro/vessel-service/proto"
	"fmt"
	"log"

	"github.com/micro/cli/v2"
	micro "github.com/micro/go-micro/v2"
)

func main() {
	fmt.Println("start main")
	server := micro.NewService(
		micro.Name("go.micro.srv.consignment"),
		micro.Version("latest"),
	)
	server.Init(
		micro.Action(func(c *cli.Context) error {
			log.Println("Info", "user-srv", "user-srv is start ...")
			db, _, err := a.InitGormDB()
			if err != nil {
				fmt.Println(err)
				// return nil, nil, nil
			}
			vClient := vesselPb.NewVesselServiceClient("go.micro.srv.vessel", server.Client())
			consiRepo := &consignment.ConsignmentPepo{
				DB: db,
			}
			consignmentSrv := &service.ConsignmentSrv{
				Repo:         consiRepo,
				vesselClient: vClient,
			}
			pb.RegisterShippingServiceHandler(server.Server(), consignmentSrv)
			return nil
		}),
		micro.AfterStop(func() error {
			log.Println("Info", "user-srv", "user-srv is stop ...")
			return nil
		}),
		micro.AfterStart(func() error {
			return nil
		}),
	)

	//启动service
	if err := server.Run(); err != nil {
		log.Panic("user-srv服务启动失败 ...")
	}
}
