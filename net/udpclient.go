package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {

	conn, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 8050,
	})
	defer conn.Close()
	if err != nil {
		fmt.Println("dial fail err,", err)
	}

	for {
		reader := bufio.NewReader(os.Stdin)
		inputstr, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("read fail ,err,", err)
			continue
		}
		input := strings.Trim(inputstr, "\n")
		_, err = conn.Write([]byte(input))
		if err != nil {
			fmt.Println("sent fail err,", err)
			continue
		}
		var buf [128]byte
		n, add, err := conn.ReadFromUDP(buf[:])
		if err != nil {
			fmt.Println("sent fail err,", add, err)
			continue
		}
		fmt.Println("accept xs:", string(buf[:n]), add)
	}
}
