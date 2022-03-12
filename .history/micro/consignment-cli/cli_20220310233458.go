package main

import (
	"context"
	pb "demo1/micro/consignment-service/proto/consignment"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"

	"github.com/micro/go-micro/v2"
)

const (
	ADDRESS  = "localhost:8080"
	FileName = "consignment.json"
)

func parseFile(fileDir string) (*pb.Consignment, error) {

	var consignment *pb.Consignment

	byte, err := ioutil.ReadFile(fileDir)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(byte, &consignment)
	if err != nil {
		return nil, errors.New("json parse fail err")
	}
	return consignment, nil
}

func main() {
	filename := FileName
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}
	p, err := parseFile(filename)
	if err != nil {
		log.Printf("parsefile fail er:%v", err)
	}
	service := micro.NewService(
		micro.Name("go.micro.srv.consignment.client"),
		micro.Version("latest"),
		micro.Address(":50051"),
	)

	service.Init()
	client := pb.NewShippingServiceClient("go.micro.srv.consignment", service.Client())
	resp, err := client.CreateConsignment(context.Background(), p)
	if err != nil {
		log.Printf("create consignment fail err %v", err)
		return
	}
	log.Printf("created is %v", resp.Created)

	resp, err = client.GetAll(context.Background(), &pb.GetRequest{})
	if err != nil {
		log.Printf("get consignments fail err %v", err)
		return
	}
	for item := range resp.Consignments {
		log.Println(item)
	}
}
