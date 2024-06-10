package main

import "fmt"

func main() {
	a := 42
	b := &a
	fmt.Println(a, *b) // Prints: 42 42
	*b = 21
	fmt.Println(a, *b) // Prints: 21 21
}
