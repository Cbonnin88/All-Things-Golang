package main

import "fmt"

func main() {
	// A function is a block of code designed to be reused
	d := 10
	e := 20
	f := 21009

	fmt.Println(tripleNum(d))
	fmt.Println(tripleNum(e))
	fmt.Println(tripleNum(f))
	jobOffer()
	max(20, 40)

}

// A simple function with parameters and a return type
func tripleNum(num int) int {
	return num * 3
}

// A Simple function with no parameters
func jobOffer() {
	fmt.Println("Congrats, you're hired")
}

// A Simple function with two parameters:
func max(num1, num2 int) {
	if num1 > num2 {
		fmt.Println(num1)
	} else {
		fmt.Println(num2)
	}
}
