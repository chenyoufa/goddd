package main

import (
	"bufio"
	"encoding/binary"
)

func Decode() (string, error) {

}
func Encode(r bufio.Reader) ([]byte, error) {
	var buteval []byte
	butevalue := r.Read(buteval)
	buf := r.Buffered()
	binary.Write(buf)

}
