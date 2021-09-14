package main

import "fmt"

func main() {
	// Block leve scope, only accessible in this function
	// Order matters !!!
	y := "You stupid man-thing!!!"
	fmt.Println(y)
}
