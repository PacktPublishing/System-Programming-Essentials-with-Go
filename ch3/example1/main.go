package main

import (
	"fmt"
	"unsafe"

	"golang.org/x/sys/unix"
)

func main() {
	// The native way to print "Hello, World!" to stdout
	fmt.Println("Hello, World!")

	// The overly complicated way to print "Hello, World!" to stdout
	unix.Syscall(unix.SYS_WRITE, 1,
		uintptr(unsafe.Pointer(&[]byte("Hello, World!")[0])),
		uintptr(len("Hello, World!")),
	)
}
