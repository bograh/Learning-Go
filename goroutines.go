package main

import (
	"fmt"
	"time"
)

// A goroutine is a lightweight thread of execution

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	// Suppose we have a function call f(s). Here's how
	// we'd call that in the usual way, running is sync.
	 f("direct")

	// To invoke this fxn in a goroutine, use `go f(s)`.
	// This new goroutine will exec concurrently with the
	// calling one.
	go f("goroutine")

	// You can also start a goroutine for an anon fxn call
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second)
	fmt.Println("DONE ✅")

	// Our two fxn calls are running async in seperate
	// goroutines now. Wait for them to finish

}