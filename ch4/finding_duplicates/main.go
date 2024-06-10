package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func computeFileHash(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}

func findDuplicateFiles(rootDir string) (map[string][]string, error) {
	duplicates := make(map[string][]string)

	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			hash, err := computeFileHash(path)
			if err != nil {
				return err
			}

			duplicates[hash] = append(duplicates[hash], path)
		}

		return nil
	})

	return duplicates, err
}

func main() {
	duplicates, err := findDuplicateFiles(".")
	if err != nil {
		fmt.Println("Error finding duplicate files:", err)
		return
	}

	for hash, files := range duplicates {
		if len(files) > 1 {
			fmt.Printf("Duplicate Hash: %s\n", hash)
			for _, file := range files {
				fmt.Println("  -", file)
			}
		}
	}
}
