package main

import "golang.org/x/sys/unix"

func main() {
	unix.Write(1, []byte("Hello, World!"))
}
