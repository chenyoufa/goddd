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

	reader := bufio.NewReader(os.Stdin)
	infoinput, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("ReadString err,", err)
	}
	info := strings.Trim(infoinput, "\n")
	n, err := conn.Write([]byte(info))
	if err != nil {
		fmt.Println("write fail err,", err)
	}

	var buf []byte
	n, err = conn.Read(buf)
	if err != nil {
		fmt.Println(err)
	}
	val := string(buf[:n])
	fmt.Println(val)

}
