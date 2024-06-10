package main

import (
	"fmt"
	"os"
)

func main() {
	info, err := os.Stat("example.txt")
	if err != nil {
		panic(err)
	}
	fmt.Printf("File name: %s\n", info.Name())
	fmt.Printf("File size: %d\n", info.Size())
	fmt.Printf("File permissions: %s\n", info.Mode())
	fmt.Printf("Last modified: %s\n", info.ModTime())
}
