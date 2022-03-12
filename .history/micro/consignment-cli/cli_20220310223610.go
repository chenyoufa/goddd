package main

import (
	pb "demo1/micro/consignment-service/proto/consignment"
	"encoding/json"
	"io/ioutil"
)

func parseFile(fileDir string) (*pb.Consignment, error) {

	var consignment *pb.Consignment

	byte, err := ioutil.ReadFile(fileDir)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(byte, consignment)
	if err != nil {

	}
}

func main() {

}
