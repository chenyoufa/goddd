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

	err := binary.Write(pkg, binary.LittleEndian, length)
	if err != nil {

	}
	err = binary.Write(pkg, binary.LittleEndian, []byte(message))
	if err != nil {

	}
}
