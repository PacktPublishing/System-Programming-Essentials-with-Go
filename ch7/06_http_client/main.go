package main

import (
	"bufio"
	"fmt"
	"net"
	"net/textproto"
)

const socketPath = "/tmp/go-server.sock"

func main() {
	conn, err := net.Dial("unix", socketPath)
	if err != nil {
		fmt.Println("Error connecting to the Unix socket:", err)
		return
	}
	defer conn.Close()

	// Make an HTTP request
	request := "GET / HTTP/1.1\r\n" +
		"Host: localhost\r\n" +
		"\r\n"
	_, err = conn.Write([]byte(request))
	if err != nil {
		fmt.Println("Error sending the request:", err)
		return
	}

	// Read the response
	reader := bufio.NewReader(conn)
	tp := textproto.NewReader(reader)

	// Read and print the status line
	statusLine, err := tp.ReadLine()
	if err != nil {
		fmt.Println("Error reading the status line:", err)
		return
	}
	fmt.Println("Status Line:", statusLine)

	// Read and print headers
	headers, err := tp.ReadMIMEHeader()
	if err != nil {
		fmt.Println("Error reading headers:", err)
		return
	}
	for key, values := range headers {
		for _, value := range values {
			fmt.Printf("%s: %s\n", key, value)
		}
	}

	// Read and print the body (assuming it's text-based)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err.Error() != "EOF" {
				fmt.Println("Error reading the response body:", err)
			}
			break
		}
		fmt.Print(line)
	}
}
