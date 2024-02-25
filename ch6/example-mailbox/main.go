package main

import (
	"fmt"
	"os"
	"sync"

	"golang.org/x/sys/unix"
)

var mailboxPath = "/tmp/task_mailbox" // Can be created by `mkfifo /tmp/task_mailbox`

func main() {
	// Check if the mailbox exists
	if !namedPipeExists(mailboxPath) {
		fmt.Println("The mailbox does not exist.")
		// Set up the mailbox (named pipe)
		fmt.Println("Creating the task mailbox...")
		if err := unix.Mkfifo(mailboxPath, 0666); err != nil {
			fmt.Println("Error setting up the task mailbox:", err)
			return
		}
	}

	// Open the named pipe for read and write
	mailbox, err := os.OpenFile(mailboxPath, os.O_RDWR, os.ModeNamedPipe)
	if err != nil {
		fmt.Println("Error opening named pipe:", err)
	}
	defer mailbox.Close()

	wg := &sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		ReadTask(mailbox)
	}()

	go func() {
		defer wg.Done()

		i := 0
		for i < 10 {
			SendTask(mailbox, fmt.Sprintf("Task %d\n", i))
			i++
		}
		// Close the mailbox
		SendTask(mailbox, "EOD\n")
		fmt.Println("All tasks sent.")
	}()

	wg.Wait()
}

func namedPipeExists(pipePath string) bool {
	_, err := os.Stat(pipePath)
	if err == nil {
		return true // The named pipe exists.
	}
	if os.IsNotExist(err) {
		return false // The named pipe does not exist.
	}
	fmt.Println("Error checking named pipe:", err)
	return false
}
