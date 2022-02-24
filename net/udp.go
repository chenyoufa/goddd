package main

import (
	"fmt"
	"net"
)

func main() {

	listen, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 8050,
	})

	if err != nil {
		fmt.Println(err)
		return
	}
	defer listen.Close()

	for {
		var buf [128]byte
		n, addr, err := listen.ReadFromUDP(buf[:])
		if err != nil {
			fmt.Println("read fail err", err)
			continue
		}
		rval := string(buf[:n])
		fmt.Println("收到客户端信息:", addr, rval)

		_, err = listen.WriteToUDP([]byte(rval+";r"), addr)
		if err != nil {
			fmt.Println("set fail ", err)
			continue
		}
	}

}
