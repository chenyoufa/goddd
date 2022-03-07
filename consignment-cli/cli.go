package main

import (
	"context"
	pb "demo1/consignment-service/proto/consignment"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"

	"google.golang.org/grpc"
)

const (
	ADDRESS = "localhost:8080"
	FILE    = "consignment.json"
)

func ParseFile(fileName string) (*pb.Consignment, error) {

	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var consignment *pb.Consignment
	err = json.Unmarshal(bytes, &consignment)
	if err != nil {
		return nil, errors.New("json unmarshal fail")
	}
	return consignment, nil
}
func main() {

	conn, err := grpc.Dial(ADDRESS, grpc.WithInsecure())
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()

	client := pb.NewShippingServerClient(conn)

	filename := FILE
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}
	consignment, err := ParseFile(filename)
	if err != nil {
		log.Printf("parse json fail ", err)
	}
	log.Println(consignment.Description)

	resp, err := client.CreateConsignment(context.Background(), consignment)
	if err != nil {
		log.Println("create consignment fail ", err)
	}
	log.Printf("created is:%v  ", resp.Created)

	// 列出目前所有托运的货物
	resp, err = client.GetConsignments(context.Background(), &pb.GetRequest{})
	if err != nil {
		log.Fatalf("failed to list consignments: %v", err)
	}

	for _, c := range resp.Consignments {
		log.Printf("%+v", c)
	}
}
