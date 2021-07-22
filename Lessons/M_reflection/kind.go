package main

import (
	"fmt"
	"reflect"
)

type animal struct {
	species		string

}
func main() { //nolint:typecheck
	// kind.() tells us what the type is made out of, kind.() also determines which methods are valid and which panic

	kindOf := reflect.TypeOf(animal{})
	fmt.Println("The name of our struct is",kindOf.Name())
	fmt.Println("Our struct is made up of",kindOf.Kind())
	fmt.Println("")


	// Elem() returns a reflect.Type
	// Elem() is valid when Kind.() is a slice, pointer, map, channel, array
	fmt.Println("Using .Elem() and .Kind() with a pointer:")
	pointer := reflect.TypeOf(&animal{})
	fmt.Println("The name of our pointer is",pointer.Elem().Name())
	fmt.Println("Our pointer is of type",pointer.Elem().Kind())






}
