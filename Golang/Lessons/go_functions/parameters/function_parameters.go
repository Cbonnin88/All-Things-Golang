package main

import "fmt"

func salesTaxPhila(salesTax, price float32) float32 {
	salesTax = price * (salesTax / 100)
	return price + salesTax
}

func main() {
	totalPrice := salesTaxPhila(8.0, 124)
	fmt.Println(totalPrice)

	myAge := 32
	myDogsAge := computeDogYears(myAge)
	fmt.Println(myDogsAge)

	totalPriceFr := vatTaxFr(123)
	fmt.Println(totalPriceFr)
}

func computeDogYears(humanYears int) int {
	dogYears := humanYears * 7
	return dogYears
}

func vatTaxFr(price float32) float32 {
	vatFr := (price / 200) * 100 * .20
	return price + vatFr
}
