package main

import (
	"bufio"
	"fmt"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()
	for {
		var buf []byte
		reader := bufio.NewReader(conn)

		n, err := reader.Read(buf)
		if err != nil {
			fmt.Println("reader err:", err)
			break
		}
		strconv := string(buf[:n])
		fmt.Println("access string :", strconv)
		conn.Write([]byte(strconv))
	}
}

func main() {

	listen, err := net.Listen("tcp", "127.0.0.1:8520")
	if err != nil {
		fmt.Println("Listen fail err:", err)
		return
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept fail err:", err)
			continue
		}
		go process(conn)
	}

}
