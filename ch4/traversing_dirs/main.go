package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	err := filepath.WalkDir(".", func(path string, d os.DirEntry, err error) error {
		if err != nil {
			fmt.Println("Error:", err)
			return nil
		}

		if d.IsDir() {
			fmt.Println("Directory:", path)
		} else {
			fmt.Println("File:", path)
		}
		return nil
	})

	if err != nil {
		fmt.Println("Error walking the path:", err)
	}
}
