package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	// Create and run the echo command
	echoCmd := exec.Command("echo", "Hello, world!")
	echoOutput, err := echoCmd.Output()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error running echoCmd: %v\n", err)
		return
	}

	// Create the grep command with the output of echoCmd as its input
	grepCmd := exec.Command("grep", "Hello")
	grepCmd.Stdin = strings.NewReader(string(echoOutput))

	// Capture the output of grepCmd
	grepOutput, err := grepCmd.Output()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error running grepCmd: %v\n", err)
		return
	}

	// Print the output of grepCmd
	fmt.Printf("Output of grep: %s", grepOutput)
}
