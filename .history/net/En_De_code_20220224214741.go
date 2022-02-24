package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
)

func Decode(r bufio.Reader) (string, error) {
	vbytes, err := r.Peek(4)
	if err != nil {
		fmt.Println("read fail err,", err)
		return "", err
	}
	buffer := bytes.NewBuffer(vbytes)
	var length int
	err = binary.Read(buffer, binary.LittleEndian, &length)
	if err != nil {
		fmt.Println("read length fail err,", err)
		return "", err
	}
	if r.Buffered()<(length+4)

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
