package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/fsnotify/fsnotify"
)

func main() {
	watchPath := "/path/to/your/directory"

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal("Error creating watcher:", err)
	}
	defer watcher.Close()

	err = watcher.Add(watchPath)
	if err != nil {
		log.Fatal("Error adding watch:", err)
	}

	go func() {
		for {
			select {
			case event := <-watcher.Events:
				// Handle the event
				fmt.Printf("Event: %s\n", event.Name)
			case err := <-watcher.Errors:
				log.Println("Error:", err)
			}
		}
	}()

	// Create a channel to receive signals
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGINT)

	// Block until a SIGINT signal is received
	<-signalCh

	fmt.Println("Received SIGINT. Exiting...")
}
