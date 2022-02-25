package main

import (
	"fmt"
	"time"
)

func main() {

	time1 := time.NewTimer(2 * time.Second)
	t1 := time.Now()
	fmt.Printf("t1:%v\n", t1)
	t2 := <-time1.C

	fmt.Printf("t2:%v\n", t2)

}
