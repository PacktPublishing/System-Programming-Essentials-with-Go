package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	path := "/home/user/documents/myfile.txt"
	dir, file := filepath.Split(path)
	fmt.Println("Directory:", dir)
	fmt.Println("File:", file)
}
