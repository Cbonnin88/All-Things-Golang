package main

import (
	"fmt"
	"reflect"
)

type person struct {
	Name	string
	Age		int
	isHuman bool
}

func main() { //nolint:typecheck

	// .valueOf() takes in any kind of variable
	pe := person{
		Name: "Chris",
		Age:   32,
		isHuman: true,
	}
	fmt.Println("The values of this struct:",reflect.ValueOf(pe))
	fmt.Println("")

	// .ValueOf() with a slice:
	slice := reflect.ValueOf([]float64{1.2,3.4,5.6,7.8,9.0})
	slice2 := reflect.ValueOf([]string{"This", "is", "a", "string","slice"})
	fmt.Println(slice, slice2)
	fmt.Println("")

	// Writing a Value:
	num := 100
	numPtrVal := reflect.ValueOf(&num)
	fmt.Println("The value of our pointer is:",numPtrVal)
	numVal := numPtrVal.Elem()
	fmt.Println("The element of our variable is:",numVal)
	fmt.Println("")

	// Creating a Value:
	var x bool
	boolType := reflect.TypeOf(x)
	fmt.Println("our varible is of type:",boolType)

	newBoolVal := reflect.New(boolType)
	fmt.Println("Our pointers address:",newBoolVal)

	newBoolVal.Elem().SetBool(false)
	newBool := newBoolVal.Elem().Bool()
	fmt.Println("We've created a new value and now its",newBool)









}
