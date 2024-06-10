package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {
	filePath := "example.txt"
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("Failed to open file: %v\n", err)
		return
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("Failed to get file info: %v\n", err)
		return
	}
	fileSize := fileInfo.Size()

	data, err := syscall.Mmap(int(file.Fd()), 0, int(fileSize), syscall.PROT_READ|syscall.PROT_WRITE, syscall.MAP_SHARED)
	if err != nil {
		fmt.Printf("Failed to mmap file: %v\n", err)
		return
	}
	defer syscall.Munmap(data)

	fmt.Printf("Initial content: %s\n", string(data))
	newContent := []byte("Hello, mmap!")
	copy(data, newContent)
	fmt.Println("Content updated successfully.")
}
