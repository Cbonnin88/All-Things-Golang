package main

import "fmt"

func main() {

	func1 := foo()
	func2, string1 := bar()

	fmt.Println(func1)
	fmt.Println(func2, string1)
}

func foo() int {
	return 1
}

func bar() (int, string) {
	return 2, "This is test"
}
