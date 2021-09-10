package main

import "fmt"

func main() {

	func() {
		fmt.Println("I am an anonymous Function")
	}()

	func(num float32) {
		fmt.Println("Now I am an anonymous function with a parameter and an argument of", num)
	}(24.4)

}
