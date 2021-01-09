package main

import "fmt"

func incomeTax(income, familyPart float32) (float32, float32) {
	familyQuotient := income / familyPart
	var incomeRate float32
	if familyQuotient > 0 && familyQuotient < 9964 {
		incomeRate = 0.0
	} else if familyQuotient >= 9965 && familyQuotient <= 25405 {
		incomeRate = 11.0
	} else if familyQuotient >= 25406 && familyQuotient <= 72643 {
		incomeRate = 30.0
	} else if familyQuotient >= 72644 && familyQuotient <= 156244 {
		incomeRate = 41.0
	} else if familyQuotient > 156245 {
		incomeRate = 45.0
	}

	return incomeRate, familyPart
}

func main() {
	var myIncome, myFamilyPart float32
	myIncome = 55000
	myFamilyPart = 3.0

	var myMoney float32
	var myFamily float32
	myMoney, myFamily = incomeTax(myIncome, myFamilyPart)

	fmt.Printf("You will pay %v percent of your taxble income this year, your family part is total if %v parts. ", myMoney, myFamily)

}
