package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"runtime/pprof"
	"time"
)

type FileInfo struct {
	Name    string
	ModTime time.Time
	Size    int64
}

func scanDirectory(dir string) (map[string]FileInfo, error) {
	results := make(map[string]FileInfo)
	err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		info, err := d.Info()
		if err != nil {
			return err
		}
		results[path] = FileInfo{
			Name:    info.Name(),
			ModTime: info.ModTime(),
			Size:    info.Size(),
		}
		return nil
	})
	return results, err
}

func compareAndEmitEvents(oldState, newState map[string]FileInfo) {
	for path, _ /* newInfo*/ := range newState {
		// ...
		go sendAlert(fmt.Sprintf("File created: %s", path))
		// ...
		go sendAlert(fmt.Sprintf("File modified: %s", path))
	}
	for path := range oldState {
		// ...
		go sendAlert(fmt.Sprintf("File deleted: %s", path))
	}
}

func sendAlert(event string) {
	fmt.Println("Alert:", event)
}

func main() {
	dirToMonitor := "./"
	interval := time.Second * 10

	f, err := os.Create("cpuprofile.out")
	if err != nil {
		// Handle error
	}
	defer f.Close()
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	currentState, err := scanDirectory(dirToMonitor)
	if err != nil {
		fmt.Println("Error scanning directory:", err)
		return
	}

	for {
		newState, err := scanDirectory(dirToMonitor)
		if err != nil {
			fmt.Println("Error scanning directory:", err)
			continue
		}
		compareAndEmitEvents(currentState, newState)
		currentState = newState
		time.Sleep(interval)
	}
}
