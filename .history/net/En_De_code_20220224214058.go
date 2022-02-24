package main

import (
	"bytes"
	"encoding/binary"
)

func Decode() (string, error) {

}
func Encode(message string) ([]byte, error) {
	length := int32(len(message))
	pkg := new(bytes.Buffer)

	binary.Write(pkg, binary.LittleEndian, length)

	binary.Write(pkg, binary.LittleEndian, []byte(message))

}
