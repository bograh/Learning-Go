package main

import (
	"errors"
	"fmt"
)


// It’s possible to use custom types as errors by 
// implementing the Error() method on them.

// A custom error type usually has the suffix “Error”.
type argError struct {
	arg int
	message string
}

// Adding this Error method makes argError implement 
// the error interface.
func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.message)
}

func f(arg int) (int, error) { 
	if arg == 42 {
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func main() {
	_, err := f(42)
	var ae *argError
	// errors.As is a more advanced versionof errors.Is.
	// It checks that the given err (or any err in its chain)
	// matches a specific err type and converts to a value
	// of that type, return true. If no match return false
	if errors.As(err, &ae) {
		fmt.Println(ae.arg)
		fmt.Println(ae.message)
	} else {
		fmt.Println("err doesn't match argError")
	}
}

