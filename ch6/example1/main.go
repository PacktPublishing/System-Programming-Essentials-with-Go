package main

import (
	"bufio"
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

	grepCmd.Stdin = pipe

	grepOut, err := grepCmd.StdoutPipe()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating StdoutPipe for grepCmd: %v\n", err)
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

	// Read and print the output of grepCmd
	scanner := bufio.NewScanner(grepOut)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := grepCmd.Wait(); err != nil {
		fmt.Fprintf(os.Stderr, "Error waiting for grepCmd: %v\n", err)
		return
	}
}
