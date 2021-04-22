package main

import "fmt"

func main() {
	fmt.Print("Enter your current account balance: ")
	var currentAccount int
	fmt.Scan(&currentAccount)

	if currentAccount > 50000 {
		fmt.Println("You are Platinum class member")
	} else if currentAccount == 50000 {
		fmt.Println("You are a gold class member")
	} else if currentAccount < 50000 && currentAccount > 20000 {
		fmt.Println("You are a sliver class member")
	} else if currentAccount <= 20000 && currentAccount > 10000 {
		fmt.Println("You are a bronze class member")
	} else if currentAccount < 10000 {
		fmt.Println("you must have at least 10000 points to be a member")
	} else {
		fmt.Println("Do you have points ?")
	}
}
