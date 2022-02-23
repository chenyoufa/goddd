package main

import (
	"fmt"
	"net"
)

func main() {

	conn, err := net.Dial("tcp", "127.0.0.1:8520")
	if err != nil {
		fmt.Println("dial fail err,", err)
	}

}
