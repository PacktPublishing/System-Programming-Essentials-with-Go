package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/fsnotify/fsnotify"
)

const (
	logFilePath = "your_log_file.log"
	maxFileSize = 1024 * 1024 * 10 // 10 MB
)

var logFile *os.File

func main() {
	var err error
	logFile, err = os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening log file:", err)
		return
	}
	defer logFile.Close()

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Println("Error creating watcher:", err)
		return
	}
	defer watcher.Close()

	err = watcher.Add(logFilePath)
	if err != nil {
		fmt.Println("Error adding log file to watcher:", err)
		return
	}

	var mu sync.Mutex

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Op&fsnotify.Write == fsnotify.Write {
					fi, err := os.Stat(logFilePath)
					if err != nil {
						fmt.Println("Error getting file info:", err)
						continue
					}
					fileSize := fi.Size()
					if fileSize >= maxFileSize {
						mu.Lock()
						rotateLogFile()
						mu.Unlock()
					}
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				fmt.Println("Error watching file:", err)
			}
		}
	}()

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGINT)

	<-signalCh
	fmt.Println("Received SIGINT. Exiting...")
}

func rotateLogFile() {
	err := closeLogFile()
	if err != nil {
		fmt.Println("Error closing log file:", err)
		return
	}

	timestamp := time.Now().Format("20060102150405")
	newLogFilePath := fmt.Sprintf("your_log_file_%s.log", timestamp)
	err = os.Rename(logFilePath, newLogFilePath)
	if err != nil {
		fmt.Println("Error renaming log file:", err)
		return
	}

	err = createLogFile()
	if err != nil {
		fmt.Println("Error creating new log file:", err)
		return
	}

	fmt.Println("Log rotated.")
}

func closeLogFile() error {
	if logFile != nil {
		return logFile.Close()
	}
	return nil
}

func createLogFile() error {
	var err error
	logFile, err = os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	log.SetOutput(logFile)
	return nil
}
