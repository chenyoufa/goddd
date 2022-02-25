package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {

	go func() {
		i := 0
		for {
			i++
			fmt.Println("new goroutine i ", i)
			time.Sleep(time.Second)
		}
	}()

	i := 0
	for {
		i++
		fmt.Println("main goroutine:i ", i)
		time.Sleep(time.Second)
		if i == 2 {
			fmt.Println("main goroutine end ")
			break
		}
	}

}
