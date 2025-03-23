package main

import "fmt"

// Channels are the pipes that connect concurrent goroutines.
// You can send values into channels from one goroutine
// and receive those values into another goroutine.

func main() {

	// Create a new channel with make(chan val-type). 
	// Channels are typed by the value they convey
	messages := make(chan string)

	// Send a val into a channel using the channel <- syntax 
	// Here we sened "ping" to the messages channel 
	// we made above, from a new goroutine
	go func() { messages <- "ping" }()

	// The <-channel syntax receives a val from the channel
	msg := <-messages
	fmt.Println(msg)
}