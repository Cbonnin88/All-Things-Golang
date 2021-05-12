package main

import "fmt"

func main() {
	fmt.Println("Your loan interest rate")

	fmt.Print("enter your current annual salary: ")
	var salary int
	_, err := fmt.Scan(&salary)
	if err != nil {
		return
	}

	if salary >= 100000 {
		fmt.Println("You are eligible for a interest free loan")
	} else if salary >= 75000 && salary < 100000 {
		fmt.Println("You are eligible for a loan at 0.25 percent interest")
	} else if salary >= 50000 && salary < 75000 {
		fmt.Println("You are eligible for a loan at 0.50 percent interest rate")
	} else if salary >= 25000 && salary < 50000 {
		fmt.Println("You are eligible for a loan at 0.75 percent interest rate")
	} else {
		fmt.Println("you are Ineligible for this loan")
	}

}
