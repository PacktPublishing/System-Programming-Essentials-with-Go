package main

import (
	"fmt"
	"os"
)

func SendTask(pipe *os.File, data string) error {
	_, err := pipe.WriteString(data)
	if err != nil {
		return fmt.Errorf("error writing to named pipe: %v", err)
	}
	return nil
}
