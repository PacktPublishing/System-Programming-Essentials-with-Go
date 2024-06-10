package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	echoCmd := exec.Command("echo", "Hello, world!")
	grepCmd := exec.Command("grep", "Hello")

	pipe, err := echoCmd.StdoutPipe()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating StdoutPipe for echoCmd: %v\n", err)
		return
	}

	if err := grepCmd.Start(); err != nil {
		fmt.Fprintf(os.Stderr, "Error starting grepCmd: %v\n", err)
		return
	}

	if err := echoCmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error running echoCmd: %v\n", err)
		return
	}

	if err := pipe.Close(); err != nil {
		fmt.Fprintf(os.Stderr, "Error closing pipe: %v\n", err)
		return
	}

	if err := grepCmd.Wait(); err != nil {
		fmt.Fprintf(os.Stderr, "Error waiting for grepCmd: %v\n", err)
		return
	}
}
