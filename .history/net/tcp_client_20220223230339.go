package main

import "net"

func main() {

	conn, err := net.Dial("tcp", "127.0.0.1:5028")

}
