package main

import (
	"fmt"
	"os"
)

func main() {
	fileInfo, err := os.Stat("example.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	permissions := fileInfo.Mode().Perm()
	permissionString := fmt.Sprintf("%o", permissions)
	fmt.Printf("Permissions: %s\n", permissionString)
}
