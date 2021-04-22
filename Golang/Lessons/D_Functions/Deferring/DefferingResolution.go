package main

import (
	"fmt"
)

func takeHomePay(income, payAsYouGo float64) float64 {
	taxRate := 0.22
	netPayBeforeWithhold := income - income*taxRate
	defer fmt.Println("Net pay before withhold :", netPayBeforeWithhold, "EUR") // our defer statement
	beforeSource := netPayBeforeWithhold - payAsYouGo
	difference := income - netPayBeforeWithhold
	fmt.Println("The difference between your Gross and net Pay is:", difference, "EUR")
	return beforeSource
}
func main() {
	fmt.Println("Total Net Pay :", takeHomePay(2000, 2.5), "EUR")
}
