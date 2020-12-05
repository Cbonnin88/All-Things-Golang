package main

import "fmt"

func main() {
	currentBankBalance := 4559
	budgetLimit := 1200

	if currentBankBalance >= 1100 && budgetLimit >= 1100 {
		fmt.Println("You have a positive account balance")
	} else {
		fmt.Println("You have a negative account balance")
	}

	// the not or ! operator negates (reverses) the value of a boolean:

	rich := true
	fmt.Println(!rich)
	poor := false
	fmt.Println(!poor)

	if !rich {
		fmt.Println("I don't need money")
	} else {
		fmt.Println("I hate being poor")
	}
	if !poor {
		fmt.Println("I need more money")
	} else {
		fmt.Println("I'm living pretty well")
	}
}
