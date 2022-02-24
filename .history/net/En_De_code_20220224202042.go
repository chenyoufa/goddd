package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
)

func Encode(message string) ([]byte, error) {

	var length = int32(len(message))
	var pkg = new(bytes.Buffer)
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

func Decode(reader *bufio.Reader) (string, error) {
	lenghtByte, _ := reader.Peek(4) // 读取前4个字节的数据

	lengthBuff := bytes.NewBuffer(lenghtByte)

	var length int32
	err := binary.Read(lengthBuff, binary.LittleEndian, &length)

	if err != nil {
		return "", err
	}
	if int32(reader.Buffered()) < length+4 {
		return "", err
	}
	pack := make([]byte, int(4+length))
	_, err = reader.Read(pack)
	if err != nil {
		return "", err
	}
	return string(pack[4:]), nil
}

func main() {

	msg := `Hello, Hello. How are you?`
	data, err := Encode(msg)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)
	lengthBuff := bytes.NewBuffer(data)

	a := bufio.NewReader(lengthBuff)
	Decode(a)
}
