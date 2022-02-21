package casbin

import (
	"fmt"
	"log"

	casbin "github.com/casbin/casbin/v2"
)

func Check(e *casbin.Enforcer, sub, obj, act string) {
	ok, _ := e.Enforce(sub, obj, act)
	if ok {
		fmt.Printf("%s CAN %s %s \n", sub, act, obj)
	} else {
		fmt.Printf("%s Cannot %s %s \n", sub, act, obj)
	}
}

func main() {
	e, err := casbin.NewEnforcer("./baseconfig/model.conf", "./baseconfig/policy.csv")
	if err != nil {
		log.Fatalf("NewEnforecer failed:%v \n", err)
	}
}
