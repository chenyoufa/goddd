package main

import (
	pb "demo1/micro/consignment-service/proto/consignment"
	"fmt"
	"time"

	"github.com/devfeel/mapper"
)

type (
	User struct {
		Name string
		Age  int
		Id   string `mapper:"_id"`
		AA   string `json:"Score"`
		Time time.Time
	}

	Student struct {
		Name  string
		Age   int
		Id    string `mapper:"_id"`
		Score string
	}

	Teacher struct {
		Name  string
		Age   int
		Id    string `mapper:"_id"`
		Level string
	}
)
type Consignment struct {
	Name        string
	Description string `gorm:"size:20;not null;"`
	Weight      int32  `gorm:"index;not null"`
	Containers  []Container
	VesselId    string `gorm:"size:20;not null"`
}
type Container struct {
	UserId        string `gorm:"size:20;not null;"`
	CoumenterId   string `gorm:"size:20;not null"`
	Origin        string `gorm:"size:20;not null"`
	ConsignmentID uint
}

func init() {
	mapper.Register(&User{})
	mapper.Register(&Student{})

	mapper.Register(&Consignment{})
	mapper.Register(&pb.Consignment{})
}

func main() {
	user := &User{}
	userMap := &User{}
	teacher := &Consignment{}
	student := &Student{Name: "test", Age: 10, Id: "testId", Score: "100"}
	valMap := make(map[string]interface{})
	valMap["Name"] = "map"
	valMap["Age"] = 10
	valMap["_id"] = "x1asd"
	valMap["Score"] = 100
	valMap["Time"] = time.Now()

	mapper.Mapper(student, user)
	mapper.AutoMapper(student, teacher)
	mapper.MapperMap(valMap, userMap)

	consign := &Consignment{}
	// pbq := &pb.Consignment{Id: "100", Description: "des", Weight: 10, VesselId: "10086"}

	// mapper.MapperMap(valMap, consign)

	// fmt.Println("user:", user)
	fmt.Println("teacher", teacher)
	// fmt.Println("userMap:", userMap)
}
