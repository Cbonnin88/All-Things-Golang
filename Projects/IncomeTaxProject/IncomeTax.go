package main

import "fmt"

func incomeTax(income, familyPart float32) (float32, float32) {
	familyQuotient := income / familyPart
	var incomeRate float32
	switch {
	case familyQuotient >= 0 && familyQuotient < 10084:
		incomeRate = 0.0
	case familyQuotient >= 10085 && familyQuotient <= 25710:
		incomeRate = 11.0
	case familyQuotient >= 25711 && familyQuotient <= 73516:
		incomeRate = 30.0
	case familyQuotient >= 73517 && familyQuotient <= 158222:
		incomeRate = 41.0
	case familyQuotient > 158222:
		incomeRate = 45.0
	default:
		fmt.Print("Error")
	}

	return incomeRate, familyPart
}

func main() {
	var myIncome, myFamilyPart float32
	fmt.Print("Enter your annual Salary: ")
	_, _ = fmt.Scan(&myIncome)
	fmt.Print("Enter your family quotient: ")
	_, _ = fmt.Scan(&myFamilyPart)
	var myMoney float32
	var myFamily float32
	myMoney, myFamily = incomeTax(myIncome, myFamilyPart)
	fmt.Printf("You will pay %v percent of your taxble income this year, your family part is total if %v parts. ", myMoney, myFamily)
}
