package main

import (
	"fmt"
	"os"
)

func main() {
	info, err := os.Stat("example.txt")
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("File does not exist")
		} else {
			panic(err)
		}
	}
	fmt.Printf("File name: %s\n", info.Name())
}
