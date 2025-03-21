package main

import "fmt"

// Go supports embedding of structs and interfaces to
// express a more seamless composition of types. This not
// to be confused with //go:embed which is a go directive
// to embed files and folders

type base struct {
	num int
}

func (b base) describe()