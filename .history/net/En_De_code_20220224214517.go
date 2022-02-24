package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
)

func Decode(r bufio.Reader) (string, error) {
	vbytes, err := r.Peek(4)
	buffer := bytes.NewBuffer(vbytes)

}
func Encode(message string) ([]byte, error) {
	length := int32(len(message))
	pkg := new(bytes.Buffer)
	err := binary.Write(pkg, binary.LittleEndian, length)
	if err != nil {
		return nil, err
	}
	err = binary.Write(pkg, binary.LittleEndian, []byte(message))
	if err != nil {
		return nil, err
	}
	return pkg.Bytes(), nil
}
