package casbin

import (
	"fmt"
	"log"
	"os"

	casbin "github.com/casbin/casbin/v2"
)

func check(e *casbin.Enforcer, sub, obj, act string) bool {
	ok, _ := e.Enforce(sub, obj, act)
	if ok {
		fmt.Printf("%s CAN %s %s \n", sub, act, obj)
	} else {
		fmt.Printf("%s Cannot %s %s \n", sub, act, obj)
	}
	return ok
}

func New() {

	str, _ := os.Getwd()
	fmt.Println(str)
	e, err := casbin.NewEnforcer("demo1/casbin/baseconfig/model.conf", "demo1/casbin/baseconfig/policy.csv")
	if err != nil {
		log.Fatalf("NewEnforecer failed:%v \n", err)
	}

	check(e, "dajun", "data1", "read")
	check(e, "dajun", "data1", "read")
	check(e, "lizi", "data2", "write")
	check(e, "dajun", "data1", "write")
	check(e, "dajun", "data2", "read")
}
