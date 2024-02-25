package main

import (
	"bufio"
	"fmt"
	"os"
)

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
