package main

import (
	"fmt"
	"time"
)

func main() {
	say("hello")
	say("world")
}
func say(s string) {
	for i := 1; i < 5; i++ {
		time.Sleep(500 * time.Millisecond)
		fmt.Println(s)
	}
}
