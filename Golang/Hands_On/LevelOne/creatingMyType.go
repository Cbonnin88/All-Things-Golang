package main

import "fmt"

func main() {
	type number int
	var x number = 2000

	fmt.Println("The value of x is", x)
	fmt.Printf("'x' is an %T type\n", x)
	x = 42
	fmt.Println("'x' has changed from 2000 to", x)

}
