package main

import (
	"fmt"
	"sync"
)

var once sync.Once

func setup() {
	fmt.Println("Initializing...")
}

func main() {
	once.Do(setup)
	once.Do(setup) // This won't execute setup again
}
