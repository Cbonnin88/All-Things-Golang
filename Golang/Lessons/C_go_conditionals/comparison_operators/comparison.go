package main

import "fmt"

/*
Comparison Operators:
== means 'equal to'
!= means 'NOT equal to'
< means 'less than'
> means 'greater than'
<= means 'less than or equal to'
>= means 'greater than or equal to'
*/

func main() {
	lotteryNumber := "13-12-99-76"
	yourGuess := "12-13-66-03"

	if lotteryNumber == yourGuess {
		fmt.Println("You won the lottery")
	} else {
		fmt.Println("Sorry, try again")
	}
	currentAccountBalance := 1100
	monthlyBudgetLimit := 1200

	if currentAccountBalance < monthlyBudgetLimit {
		fmt.Println("You are approaching your budget limit for the month")
	} else {
		fmt.Println("Your current balance is 4500 euros")
	}

	var compare = 100
	var compare2 = 200
	fmt.Println("100 is not greater than or equal to 200, so this will give us", compare >= compare2)

	num1 := 200
	num2 := 200
	fmt.Println("200 is equal to 200, so this will give us", num1 <= num2)

}
