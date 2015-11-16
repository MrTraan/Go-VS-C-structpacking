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

	for i := range input {
		input[i].id = int32(i)
		input[i].comment[0] = 116
		input[i].comment[1] = 101
		input[i].comment[2] = 115
		input[i].comment[3] = 116
		input[i].comment[4] = 0
	}

	//file, err := os.Open("data")
	file, err := os.OpenFile("datago", os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0666)
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
