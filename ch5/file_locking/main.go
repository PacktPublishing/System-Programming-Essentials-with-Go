package main

import (
	"fmt"
	"io"
	"os"
	"syscall"
)

func main() {
	file, err := os.Open("yourfile.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	lock := syscall.Flock_t{
		Type:   syscall.F_WRLCK, // Lock type (F_RDLCK for read lock, F_WRLCK for write lock)
		Whence: io.SeekStart,    // Offset base (0 for the start of the file)
		Start:  0,               // Start offset
		Len:    0,               // Length of the locked region (0 for entire file)
	}

	if err := syscall.FcntlFlock(file.Fd(), syscall.F_SETLK, &lock); err != nil {
		fmt.Println("Error locking file:", err)
		return
	}

	// Your file operations here

	lock.Type = syscall.F_UNLCK
	if err := syscall.FcntlFlock(file.Fd(), syscall.F_SETLK, &lock); err != nil {
		fmt.Println("Error unlocking file:", err)
	}
}
