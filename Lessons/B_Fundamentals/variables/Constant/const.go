package main

import (
	"fmt"
)

func main() {
	// Constant use "const" as a keyword
	// Constant must be initialized when you set them:

	// Example of an integer constants:
	const a = 12
	fmt.Println("our example of an integer constance:", a)
	// Example of a string constance:
	const b = "string constance"
	fmt.Println("This is a", b)
	// Example of a float constance
	const dec = 23.323
	fmt.Println("This is an example of a float constance:", dec)

	// Examples of type constants:

	const (
		word string  = "Christopher Bonnin"
		age  int     = 32
		avg  float64 = 3.8
	)
	fmt.Println("Type constants:")
	fmt.Printf("My name is %v and I am %v years old, my GPA is %v for this year\n", word, age, avg)

	// unlike variables, constants cannot be changed, they are locked in for life

}
