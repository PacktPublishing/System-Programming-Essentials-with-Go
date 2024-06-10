package main

import (
	"fmt"
	"os"
	"os/signal"
)

func main() {
	signals := make(chan os.Signal, 1)
	done := make(chan struct{}, 1)

	signal.Notify(signals, os.Interrupt)

	go func() {
		for {
			s := <-signals
			switch s {
			case os.Interrupt:
				fmt.Println("INTERRUPT")
				done <- struct{}{}
			default:
				fmt.Println("OTHER")
			}
		}
	}()

	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")
}
