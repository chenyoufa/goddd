package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
)

func Decode() (string, error) {

}
func Encode(r bufio.Reader) ([]byte, error) {
	var buteval []byte
	n, err := r.Read(buteval)
	if err != nil {
		fmt.Println(err)
	}
	buf := r.Buffered()
	binary.Write(buf)

}
