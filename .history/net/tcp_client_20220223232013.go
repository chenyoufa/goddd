package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {

	conn, err := net.Dial("tcp", "127.0.0.1:8520")
	if err != nil {
		fmt.Println("dial fail err,", err)
	}
	defer conn.Close()
	var buf []byte
	reader := bufio.NewReader(os.Stdin)
	infoinput, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("ReadString err,", err)
	}
	info := strings.Trim(infoinput, "\n\t")
	n, err := conn.Write([]byte(info))

}
