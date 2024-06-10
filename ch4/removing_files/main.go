package main

import (
	"fmt"
	"os"
)

func main() {
	filePath := "/path/to/your/file-or-symlink.txt"

	err := os.Remove(filePath)
	if err != nil {
		fmt.Printf("Error removing the file: %v\n", err)
		return
	}
	fmt.Printf("File removed: %s\n", filePath)
}
