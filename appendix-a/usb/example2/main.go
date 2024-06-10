package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// read fist parameter as path. e.g. /dev/sdc
	path := os.Args[1]

	if !strings.HasPrefix(path, "/dev/") {
		fmt.Println("Path must start with /dev/")
		return
	}

	file, err := os.Open("/proc/mounts")
	if err != nil {
		fmt.Printf("Error opening /proc/mounts: %v\n", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		if len(fields) >= 2 {
			device := fields[0]
			mountPoint := fields[1]
			if strings.HasPrefix(device, path) {
				mountPoint = strings.ReplaceAll(mountPoint, "\\040", " ")
				fmt.Printf("Device: %s is mounted on: %s\n", device, mountPoint)
				fmt.Println("Files:")
				err := filepath.Walk(mountPoint, func(path string, info os.FileInfo, err error) error {
					if err != nil {
						return filepath.SkipDir
					}
					fmt.Println(path)
					return nil
				})
				if err != nil {
					fmt.Printf("Error walking the path %v: %v\n", mountPoint, err)
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading /proc/mounts: %v\n", err)
	}
}
