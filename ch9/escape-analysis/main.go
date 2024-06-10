package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	p := createPerson()
	fmt.Println(p)
}

//go:noinline
func createPerson() *person {
	p := person{name: "Alex Rios", age: 99}
	return &p
}
