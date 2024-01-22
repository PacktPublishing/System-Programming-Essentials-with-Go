package example1

import (
	"os"
	"path/filepath"
	"testing"
)

// helper function to create a temporary file with a given extension
func createTempFileWithExt(dir, ext string) (string, error) {
	file, err := os.CreateTemp(dir, "*"+ext)
	if err != nil {
		return "", err
	}
	file.Close()
	return file.Name(), nil
}

func TestOrganizeFiles(t *testing.T) {
	// Create a temporary directory
	tempDir, err := os.MkdirTemp("", "organizeFilesTest")
	if err != nil {
		t.Fatalf("Failed to create temporary directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Map to store original and expected new locations of files
	expectedLocations := make(map[string]string)

	// Create temporary files and store their expected new locations
	txtFile, err := createTempFileWithExt(tempDir, ".txt")
	if err != nil {
		t.Fatalf("Failed to create temporary .txt file: %v", err)
	}
	expectedLocations[txtFile] = filepath.Join(tempDir, "txt", filepath.Base(txtFile))

	jpgFile, err := createTempFileWithExt(tempDir, ".jpg")
	if err != nil {
		t.Fatalf("Failed to create temporary .jpg file: %v", err)
	}
	expectedLocations[jpgFile] = filepath.Join(tempDir, "jpg", filepath.Base(jpgFile))

	// Run the organizeFiles function
	paths := []string{tempDir}
	events, err := organizeFiles(paths)
	if err != nil {
		t.Errorf("organizeFiles returned an unexpected error: %v", err)
	}

	// Check if the expected number of events were recorded
	if len(events) != len(expectedLocations) {
		t.Errorf("Expected %d events, got %d", len(expectedLocations), len(events))
	}

	// Check if files are moved to the correct new locations
	for original, newLocation := range expectedLocations {
		if _, err := os.Stat(newLocation); os.IsNotExist(err) {
			t.Errorf("File was not moved to expected location: %s", newLocation)
		}

		if _, err := os.Stat(original); !os.IsNotExist(err) {
			t.Errorf("Original file still exists at: %s", original)
		}
	}
}

func TestOrganizeFilesWithInvalidPath(t *testing.T) {
	paths := []string{"/path/does/not/exist"}
	_, err := organizeFiles(paths)
	if err == nil {
		t.Errorf("Expected an error for invalid path, but got none")
	}
}

func TestOrganizeFilesEmptyPaths(t *testing.T) {
	var paths []string
	events, err := organizeFiles(paths)
	if err != nil {
		t.Errorf("organizeFiles returned an unexpected error: %v", err)
	}
	if len(events) != 0 {
		t.Errorf("Expected 0 events for empty paths, got %d", len(events))
	}
}
