package main

import "fmt"

func performAddition() {
	x := 5
	y := 10

	fmt.Println("the sum of", x, "and", y, "is", x+y)
}

func main() {
	performAddition()
}

// This prints an error because the function is in a different scope that defined in the main
