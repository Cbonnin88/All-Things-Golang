package main

import "fmt"

func main() { //nolint:typecheck

	a := increment()
	b := increment()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())

	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())


}
/*
A closure is a special type of anonymous function that references variables declared outside of the
function itself. It's similar to acessing global variables which are available before the declaration of the
function
*/

func increment() func() float32 {
		var num float32
	return func() float32{
				num++
				return num
	}
}
