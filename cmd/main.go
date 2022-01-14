package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	var v interface{}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 10; i++ {
		v = i
		if (r.Intn(100) % 2) == 0 {
			v = "hello"
		}

		if b, a := v.(int); a {
			fmt.Println("a:", b, a)

		}
	}
}
