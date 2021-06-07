package main

import (
	"fmt"
	"reflect"
)

/*
	 Reflection is the ability to examine types at runtime
	It allows you to examine, modify, and create:

		variables
		functions
		structs
*/

type empty struct {}

func main() { //nolint:typecheck
	// Finding the name with .Name():
	var x int
	fmt.Println("Finding our variable type:")
	xt := reflect.TypeOf(x) // here we return the reflection Type
	fmt.Println(xt.Name()) // here we will get the name of our type
	fmt.Println("")

	// using .Name() with a function:
	fmt.Println("Finding our struct Type:")
	s := empty{}
	st := reflect.TypeOf(s)
	fmt.Println(st.Name()) // Here it returns our types name

	// Using .Name() with a pointer:
	p := reflect.TypeOf(&x)
	fmt.Println(p.Name())
}
