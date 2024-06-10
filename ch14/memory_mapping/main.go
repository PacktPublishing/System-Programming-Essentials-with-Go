package main

import (
	"fmt"

	"golang.org/x/exp/mmap"
)

func main() {
	const filename = "example.txt"

	reader, err := mmap.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer reader.Close()

	fileSize := reader.Len()
	data := make([]byte, fileSize)
	_, err = reader.ReadAt(data, 0)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	lastByte := data[fileSize-1]
	fmt.Printf("Last byte of the file: %v\n", lastByte)
}
