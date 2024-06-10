package main

import (
	"fmt"
	"os"
)

func main() {
	sourcePath := "/home/user/Documents/important_document.txt"
	symlinkPath := "/home/user/Desktop/shortcut_to_document.txt"

	err := os.Symlink(sourcePath, symlinkPath)
	if err != nil {
		fmt.Printf("Error creating symlink: %v\n", err)
		return
	}

	fmt.Printf("Symlink created: %s -> %s\n", symlinkPath, sourcePath)
}
