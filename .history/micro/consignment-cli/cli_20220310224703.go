package main

import (
	"context"
	pb "demo1/micro/consignment-service/proto/consignment"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"

	client "github.com/micro/go-micro/v2/client"
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
	err = json.Unmarshal(byte, consignment)
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

	client := pb.NewShippingServiceClient("go.micro.srv.consignment", client.DefaultClient)
	resp, err := client.CreateConsignment(context.Background(), p)
	if err != nil {
		log.Printf("create consignment fail err %v", err)
		return
	}
	log.Printf("created is %v", resp.Created)

}
