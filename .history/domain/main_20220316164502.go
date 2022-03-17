package main

import (
	"fmt"
	"log"

	"github.com/micro/cli"
	micro "github.com/micro/go-micro/v2"
)

func main() {
	fmt.Println("start main")
	server := micro.NewService(
		micro.Name("go.micro.srv.consignment"),
		micro.Version("latest"),
	)
	server.Init(
		micro.Action(func(c *cli.Context) error {}),
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
