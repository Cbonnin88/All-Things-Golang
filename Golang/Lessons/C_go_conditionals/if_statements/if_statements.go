package main

import "fmt"

func main() {
	// Simple IF statement:
	correctNumber := true
	if correctNumber {
		fmt.Println("We won the lottery !!!")
	}

	// Simple If-else statement:
	notCorrectNumber := false
	if notCorrectNumber {
		fmt.Println("You won the lottery")
	} else {
		fmt.Println("Sorry, try again ")
	}

}
