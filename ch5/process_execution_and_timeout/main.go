package main

import (
	"context"
	"fmt"
	"os/exec"
	"time"
)

func main() {
	cmd := exec.Command("sleep", "2") // Replace "sleep" "2" with your command and arguments
	timeout := 3 * time.Second        // Set your timeout duration

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	if err := cmd.Start(); err != nil {
		fmt.Println("Error starting command:", err)
		return
	}

	done := make(chan error, 1)
	go func() {
		done <- cmd.Wait()
	}()

	select {
	case <-ctx.Done():
		if err := cmd.Process.Kill(); err != nil {
			fmt.Println("Failed to kill process:", err)
		}
		fmt.Println("Process killed as timeout reached")
	case err := <-done:
		if err != nil {
			fmt.Println("Process finished with error:", err)
		} else {
			fmt.Println("Process finished successfully")
		}
	}
}
