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
}

func main() {
	var output [10]foo

	file, err := os.Open("data")
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
	err = binary.Read(reader, binary.LittleEndian, &output)
	if err != nil {
		fmt.Println("binary.Read failed:", err)
	}

	for i := range output {
		fmt.Printf("id: %d\tcomment: %s\n", output[i].id, output[i].comment)
	}
}
