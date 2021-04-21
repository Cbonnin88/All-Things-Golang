package main

import "fmt"

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

	// unlike variables, constants cannot be changed, they are locked in for life
}
