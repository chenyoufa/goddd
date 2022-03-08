package main

import (
	"context"
	pb "demo1/consignment-service/proto/consignment"
	"encoding/json"
	"errors"
	"log"
	"os"

	"google.golang.org/grpc"
)

const (
	ADDRESS = "127.0.0.1:8080"
	FileDIr = "consignment.json"
)

func PareFile(filedir string) (*pb.Consignment, error) {
	var consignment *pb.Consignment
	bytes, err := os.ReadFile(filedir)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bytes, &consignment)
	if err != nil {
		return nil, errors.New("json unmarshal fail ")
	}
	return consignment, nil
}

func main() {

	fileinfo := FileDIr
	if len(os.Args) > 1 {
		fileinfo = os.Args[1]
	}
	consignmentInfo, err := PareFile(fileinfo)
	if err != nil {
		log.Println("parsefile fail ")
	}
	conn, err := grpc.Dial(ADDRESS, grpc.WithInsecure())
	if err != nil {
		log.Println(err)
	}
	client := pb.NewShippingServiceClient(conn)
	res, err := client.CreateConsignment(context.Background(), consignmentInfo)
	if err != nil {
		log.Fatalf("create consignment fail err %v", err)
	}
	log.Printf("get created is %v", res.Created)

	res, err = client.GetAll(context.Background(), &pb.GetRequest{})
	if err != nil {
		log.Fatalf("GetAll consignment fail err %v", err)
	}
	for _, item := range res.Consignments {
		log.Println(item)
	}
}
