package main

import (
	"fmt"
	"time"
)

// Go's select lets you wait on multiple channel operations.
// Combining goroutines and channels with select is
// powerful feature of Go

func main() {
	
	c1 := make(chan string)
	c2 := make(chan string)

	// Each channel will receive a value after some amt of
	// time, to simulate e.g blocking RPC operations exec in
	// concurrent goroutines.
	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	// We'll use select to await both of these values
	// simultaneously, printing each one as it arrives.
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <- c1:
			fmt.Println("received", msg1)
		case msg2 := <- c2:
			fmt.Println("received", msg2)
		}
	}

}