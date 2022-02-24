package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func Decode() (string, error) {

}
func Encode(message string) ([]byte, error) {
	length := int32(string(message))
	pkg := new(bytes.Buffer)

	var buteval []byte
	n, err := r.Read(buteval)
	r.Buffered()
	if err != nil {
		fmt.Println(err)
	}

	binary.Write(buf)

}
