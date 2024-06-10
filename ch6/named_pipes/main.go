package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"

	"golang.org/x/sys/unix"
)

func main() {
	mailboxPath := "/tmp/task_mailbox"
	if !namedPipeExists(mailboxPath) {
		fmt.Println("The mailbox does not exist.")
		fmt.Println("Creating the task mailbox...")
		if err := unix.Mkfifo(mailboxPath, 0666); err != nil {
			fmt.Println("Error setting up the task mailbox:", err)
			return
		}
	}

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
		for i := 0; i < 10; i++ {
			SendTask(mailbox, fmt.Sprintf("Task %d\n", i))
		}
		SendTask(mailbox, "EOD\n")
		fmt.Println("All tasks sent.")
	}()

	wg.Wait()
}

func SendTask(pipe *os.File, data string) error {
	_, err := pipe.WriteString(data)
	if err != nil {
		return fmt.Errorf("error writing to named pipe: %v", err)
	}
	return nil
}

func ReadTask(pipe *os.File) error {
	fmt.Println("Reading tasks from the mailbox...")

	scanner := bufio.NewScanner(pipe)
	for scanner.Scan() {
		task := scanner.Text()
		fmt.Printf("Processing task: %s\n", task)
		if task == "EOD" {
			break
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading tasks from the mailbox: %v", err)
	}
	fmt.Println("All tasks processed.")
	return nil
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
