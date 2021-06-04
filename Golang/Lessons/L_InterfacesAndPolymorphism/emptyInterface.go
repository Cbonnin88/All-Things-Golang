package main

import "fmt"

func main() { //nolint:typecheck

	// An empty interface may hold values of any type
	// Empty interfaces are used by code that handles values of unknown type

	var i interface{}
	describe(i)

	i = 12.12
	describe(i)

	i = false
	describe(i)

	var o interface{}
	fmt.Println(o)


}

func describe(i interface{}){
	fmt.Printf("(%v, %T)\n", i, i )
}
