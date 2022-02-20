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

func New() bool {

	str, _ := os.Getwd()
	fmt.Println(str)
	e, err := casbin.NewEnforcer("../casbin/baseconfig/model.conf", "../casbin/baseconfig/policy.csv")
	// p, dajun, data1, read
	// p, lizi, data2, write
	if err != nil {
		log.Fatalf("NewEnforecer failed:%v \n", err)
	}

	if ok := check(e, "dajun", "data1", "read"); ok {
		return true
	}
	if ok := check(e, "lizi", "data2", "write"); ok {
		return true
	}
	if ok := check(e, "dajun", "data1", "write"); !ok {
		return true
	}
	if ok := check(e, "dajun", "data2", "read"); !ok {
		return true
	}
	return false
}
