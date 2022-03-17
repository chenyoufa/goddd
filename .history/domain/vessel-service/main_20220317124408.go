package main

import (
	pb "demo1/domain/vessel-service/proto"
)

func main() {

	vessels := []*pb.Vessel{
		{Id: "vessel001", Name: "Boaty McBoatface", MaxWeight: 200000, Capacity: 500},
	}

}
