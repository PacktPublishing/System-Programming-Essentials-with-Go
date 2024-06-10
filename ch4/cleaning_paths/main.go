package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	uncleanPath := "/home/user/../documents/file.txt"
	cleanPath := filepath.Clean(uncleanPath)
	fmt.Println("Cleaned path:", cleanPath)
}
