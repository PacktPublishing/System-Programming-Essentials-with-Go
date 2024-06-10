package main

import (
	"arena"
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	// Create a new arena
	mem := arena.NewArena()

	// Allocate a new reference for a Person struct in the arena
	p := arena.New[Person](mem)
	p.Name = "John Doe"
	p.Age = 30
	fmt.Printf("Person in arena: %+v\n", *p)

	// Create a slice with a predetermined capacity in the arena
	slice := arena.MakeSlice[string](mem, 100, 100)
	slice[0] = "Hello"
	slice[1] = "World"
	fmt.Printf("Slice in arena: %v\n", slice[:2])

	// Clone an object from the arena to the heap
	p2 := arena.Clone(p)
	fmt.Printf("Cloned person from arena to heap: %+v\n", *p2)

	// Free the arena to deallocate all objects at once
	mem.Free()

	// Uncommenting the following line will cause a runtime error because the arena is freed
	// p.Age = 31 // <- this is a problem

	// Demonstrate using the address sanitizer (run with `go run -asan main.go`)
	// Uncomment the below lines to see address sanitizer in action
	/*
	   o := arena.New[Person](mem)
	   mem.Free()
	   o.Age = 123 // <- this is a problem
	*/
}
