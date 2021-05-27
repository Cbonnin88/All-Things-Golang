package main

import (
	"fmt"
)

func main() {
	fmt.Println("Net pay after withhold :", takeHomePay(2000, 1.3), "EUR")

}

func takeHomePay(income, payAsYouGo float32) float32 {
	socialCharges := float32(0.22)
	netPayBeforeWithhold := income - (income*socialCharges)
	defer fmt.Println("Net pay before withhold :", netPayBeforeWithhold, "EUR") // our defer statement
	source := netPayBeforeWithhold * (float32(100) - payAsYouGo)
	difference := income - netPayBeforeWithhold
	fmt.Println("The difference between your Gross and net Pay is:", difference, "EUR")
	return source
}