package main

import "fmt"

func main() {
	star := "Polaris"
	starAddress := &star
	*starAddress = "Sirius"
	fmt.Println("We have changed the value of star from Polaris to", star, "using dereferencing.")
	// Example 2:
	legalAge := 18
	newLegalAge := &legalAge
	*newLegalAge = 21
	fmt.Println("We have changed our legal age from 18 to", legalAge)
	// Example 3:
	toChange := "Java"
	fmt.Println("When I started my training, I was learning", toChange)
	changeValue(&toChange)
	fmt.Println("and now I am learning", toChange)
}

// Using a function:
func changeValue(str *string) {
	*str = "Go"
}
func changeValue2(str string) {
	str = "changed"
}
