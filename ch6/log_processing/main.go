package main

import (
	"bufio"
	"io"
	"os"
	"strings"
	"syscall"
	"time"
)

func filterLogs(reader io.Reader, writer io.Writer) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		logEntry := scanner.Text()
		if strings.Contains(logEntry, "ERROR") {
			writer.Write([]byte(logEntry + "\n"))
		}
	}
}

func main() {
	pipePath := "/tmp/my_log_pipe"
	if err := os.RemoveAll(pipePath); err != nil {
		panic(err)
	}
	if err := syscall.Mkfifo(pipePath, 0600); err != nil {
		panic(err)
	}
	defer os.RemoveAll(pipePath)

	pipeFile, err := os.OpenFile(pipePath, os.O_RDONLY|os.O_CREATE, os.ModeNamedPipe)
	if err != nil {
		panic(err)
	}
	defer pipeFile.Close()

	go func() {
		writer, err := os.OpenFile(pipePath, os.O_WRONLY, os.ModeNamedPipe)
		if err != nil {
			panic(err)
		}
		defer writer.Close()

		for {
			writer.WriteString("INFO: All systems operational\n")
			writer.WriteString("ERROR: An error occurred\n")
			time.Sleep(1 * time.Second)
		}
	}()

	filterLogs(pipeFile, os.Stdout)
}
