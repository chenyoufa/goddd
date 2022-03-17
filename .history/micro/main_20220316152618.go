package main

import (
	"fmt"
	"time"

	"github.com/devfeel/mapper"
	"github.com/jinzhu/copier"
)

type (
	User struct {
		Name        string
		Age         int
		Id          string `mapper:"_id"`
		AA          string `json:"Score"`
		Time        time.Time
		Description string
	}

	Student struct {
		Name        string
		Age         int
		Id          string `mapper:"_id"`
		Score       string
		Description string
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
type Consignmentpro struct {
	Id          string       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Description string       `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Weight      int32        `protobuf:"varint,3,opt,name=weight,proto3" json:"weight,omitempty"`
	Containers  []*Container `protobuf:"bytes,4,rep,name=containers,proto3" json:"containers,omitempty"`
	VesselId    string       `protobuf:"bytes,5,opt,name=vessel_id,json=vesselId,proto3" json:"vessel_id,omitempty"`
	// XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	// XXX_unrecognized     []byte       `json:"-"`
	// XXX_sizecache        int32        `json:"-"`
}

func init() {
	mapper.Register(&User{})
	mapper.Register(&Student{})

	mapper.Register(&Consignment{})
	mapper.Register(&Consignmentpro{})
}

func main() {
	user := &User{}
	userMap := &User{}
	teacher := &Teacher{}
	student := &Student{Name: "nametest", Age: 10, Id: "testId", Score: "100", Description: "des"}
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
	pbq := &Consignmentpro{Id: "100", Description: "des", Weight: 10, VesselId: "10086"}

	copier.Copy(consign, pbq)
	// err := mapper.Mapper(pbq, consign)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// mapper.MapperMap(valMap, consign)

	// fmt.Println("user:", user)
	fmt.Println("consign", consign)
	// fmt.Println("userMap:", userMap)
}
