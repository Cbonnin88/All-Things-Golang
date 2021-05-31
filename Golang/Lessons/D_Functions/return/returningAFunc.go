package main

import "fmt"

func main() { //nolint:typecheck
	returnFunc := Return()
	fmt.Println(returnFunc)

	returnFunc2 := Return2()
	fmt.Printf("This func is of type:%T\n",returnFunc2)

	// Here we assign '2000' to 'x':
	x := returnFunc2()
	fmt.Println("I assigned 2000 to x, so it returns:",x)

	// We can a function this way
	fmt.Println("shorthand way of writing our return:",Return3()())

}

// A simple function that returns a string
func Return() string {
	s := "I am returning a string"
	return s
}

// We are returning a func int:
func Return2() func() int {
	return func() int {
		return 2000
	}
}

func Return3() func() float32 {
	return func() float32 {
		return 23.45
	}
}

