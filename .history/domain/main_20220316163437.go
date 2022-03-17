package main

import (
	a "demo1/domain/consignmentService"
	"demo1/domain/consignmentService/dao/consignment"
	"demo1/domain/consignmentService/service"
	"fmt"
	"log"

	pb "demo1/domain/consignmentService/proto"

	"github.com/micro/cli"
	micro "github.com/micro/go-micro/v2"
	"go.uber.org/zap"
	"gorm.io/gorm/logger"
)

func main() {
	fmt.Println("start main")
	server := micro.NewService(
		micro.Name("go.micro.srv.consignment"),
		micro.Version("latest"),
	)
	server.Init(
		micro.Action(func(c *cli.Context) {
			logger.Info("Info", zap.Any("user-srv", "user-srv is start ..."))

			db, _, err := a.InitGormDB()
			if err != nil {
				fmt.Println(err)
				// return nil, nil, nil
			}
			consiRepo := &consignment.ConsignmentPepo{
				DB: db,
			}
			consignmentSrv := &service.ConsignmentSrv{
				Repo: consiRepo,
			}

		}),
		micro.AfterStop(func() error {
			logger.Info("Info", zap.Any("user-srv", "user-srv is stop ..."))
			return nil
		}),
		micro.AfterStart(func() error {
			return nil
		}),
	)

	pb.RegisterShippingServiceHandler(server.Server(), consignmentSrv)

	if err := server.Run(); err != nil {
		log.Fatalf("failed to serve:%v", err)
	}
}
