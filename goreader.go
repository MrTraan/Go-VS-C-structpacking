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
	var output [10]foo

	if len(os.Args) != 2 {
		fmt.Println("Need file path")
		return
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("error when opening file:", err)
		return
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		fmt.Println("Couldn't get file size:", err)
		return
	}
	data := make([]byte, stat.Size())
	file.Read(data)

	reader := bytes.NewReader(data)
	for i := range output {
		err = binary.Read(reader, binary.LittleEndian, &output[i].id)
		if err != nil {
			fmt.Println("binary.Read failed:", err)
			return
		}
		err = binary.Read(reader, binary.LittleEndian, &output[i].comment)
		if err != nil {
			fmt.Println("binary.Read failed:", err)
			return
		}
		err = binary.Read(reader, binary.LittleEndian, &output[i].padding)
		if err != nil {
			fmt.Println("binary.Read failed:", err)
			return
		}
	}

	for i := range output {
		fmt.Printf("id: %d\tcomment: %s\n", output[i].id, output[i].comment)
	}
}
