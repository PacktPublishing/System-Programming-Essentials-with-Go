package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	// Start a new process
	cmd := exec.Command("ls", "-l")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
		return
	}
	// Get the current process ID
	pid := os.Getpid()
	fmt.Println("Current process ID:", pid)
}
