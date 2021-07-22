package main

import "fmt"

func main() {
	// Simple IF statement:
	correctNumber := true
	if correctNumber {
		fmt.Println("We won the lottery !!!")
	}

	// Simple If statement w/ init statement added:
	if num := 33; num == 33 {
		fmt.Println("001")
	}
	fmt.Println("Go uses ';' to")
	fmt.Println("separate statements")

	// Simple If-else statement:
	notCorrectNumber := true
	if notCorrectNumber != true {
		fmt.Println("You won the lottery")
	} else {
		fmt.Println("Sorry, try again ")
	}

}
