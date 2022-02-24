package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {

	conn, err := net.Dial("tcp", "127.0.0.1:8085")
	if err != nil {
		fmt.Println("dial fail err,", err)
		return
	}
	defer conn.Close()
	reader := bufio.NewReader(os.Stdin)

	for {

		inputinf, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			continue
		}
		input := strings.Trim(inputinf, "\n")
		conn.Write([]byte(input))
		var buf [128]byte
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Println(err)
			continue
		}
		val := string(buf[:n])
		fmt.Println("accecp val,", val)

	}

}
