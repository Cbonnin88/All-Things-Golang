package main

import "fmt"

func main() { //nolint:typecheck
	// Anonymous functions have to run in the func main():
	func() {
		fmt.Println("I am an anonymous function")
	}()
	// passing in an argument in our anonymous function
	func(x int){
		fmt.Println("I am an anonymous function with a parameter, and an argument:",x)
	}(12) // here we are passing in one argument
	doSomething()
}
func doSomething() {
	fmt.Println("I am NOT an anonymous function")
}


