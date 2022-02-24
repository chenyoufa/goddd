package main

import (
	"bufio"
	"fmt"
	"net"
)

func proccess(conn net.Conn) {
	defer conn.Close()

	for {
		var bytebuff [128]byte
		reader := bufio.NewReader(conn)
		n, err := reader.Read(bytebuff[:])
		if err != nil {
			fmt.Println(err)
			break
		}
		val := string(bytebuff[:n])
		fmt.Println("accept val ,", val)
		val = val + ";r"
		conn.Write([]byte(val))

	}
}

func main() {

	listen, err := net.Listen("tcp", "127.0.0.1:8085")
	if err != nil {
		fmt.Println("listen faile err,", err)
		return
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen faile err,", err)
			continue
		}
		go proccess(conn)
	}

}
