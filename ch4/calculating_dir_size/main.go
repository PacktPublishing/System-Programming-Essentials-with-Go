package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func calculateDirSize(path string) (int64, error) {
	var size int64

	err := filepath.Walk(path, func(filePath string, fileInfo os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !fileInfo.IsDir() {
			size += fileInfo.Size()
		}
		return nil
	})

	if err != nil {
		return 0, err
	}

	return size, nil
}

func main() {
	dirSize, err := calculateDirSize(".")
	if err != nil {
		fmt.Println("Error calculating directory size:", err)
		return
	}
	fmt.Printf("Directory size: %d bytes\n", dirSize)
}
