package main

import (
	pb "demo1/micro/consignment-service/proto/consignment"
	"encoding/json"
	"errors"
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
		return nil, errors.New("json parse fail err")
	}
	return consignment, nil
}

func main() {

}
