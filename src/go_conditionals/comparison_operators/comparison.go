package main

import "fmt"

func main() {
	lotteryNumber := "13-12-99-76"
	yourGuess := "12-13-66-03"

	if lotteryNumber == yourGuess {
		fmt.Println("You won the lottery")
	} else {
		fmt.Println("Sorry, try again")
	}
	currentAccountBalance := 4500
	monthlyBudgetLimit := 1200

	if currentAccountBalance < monthlyBudgetLimit {
		fmt.Println("You are approaching your budget limit for the month")
	} else {
		fmt.Println("Your current balance is 4500 euros")
	}

}
