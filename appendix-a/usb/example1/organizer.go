package example1

import (
	"fmt"
	"os"
	"path/filepath"
)

func organizeFiles(paths []string) ([]string, error) {
	var err error
	events := make([]string, 0)
	for _, path := range paths {
		err := filepath.WalkDir(path, func(path string, dir os.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if !dir.IsDir() {
				ext := filepath.Ext(path)
				destDir := filepath.Join(filepath.Dir(path), ext[1:]) // Remove the leading dot from the extension
				destPath := filepath.Join(destDir, dir.Name())

				// Create the destination directory if it doesn't exist
				if err := os.MkdirAll(destDir, os.ModePerm); err != nil {
					return err
				}

				// Move the file to the destination
				if err := os.Rename(path, destPath); err != nil {
					return err
				}
				events = append(events, fmt.Sprintf("Moved %s to %s\n", path, destPath))
			}
			return nil
		})

		if err != nil {
			fmt.Printf("Error walking the path %v: %v\n", path, err)
		}
	}
	return events, err
}
