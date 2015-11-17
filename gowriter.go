package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
)

type foo struct {
	id      int32
	comment [5]byte
	padding [3]byte
}

func main() {
	var input [10]foo

	if len(os.Args) != 2 {
		fmt.Println("Need file path")
		return
	}

	for i := range input {
		input[i].id = int32(i)
		input[i].comment[0] = 't'
		input[i].comment[1] = 'e'
		input[i].comment[2] = 's'
		input[i].comment[3] = 't'
		input[i].comment[4] = 0
	}

	file, err := os.OpenFile(os.Args[1], os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("error when opening file:", err)
		return
	}
	defer file.Close()

	buf := new(bytes.Buffer)
	err = binary.Write(buf, binary.LittleEndian, &input)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
		return
	}

	_, err = file.Write(buf.Bytes())
	if err != nil {
		fmt.Println("file.Write failed:", err)
		return
	}
}
