package main

import "fmt"

func main() {
	totalPrice := salesTaxPhila(8.0, 500)
	fmt.Println("The price of an new gaming console in Philadelphia with the sales tax added:",totalPrice)

	myAge := 32
	myDogsAge := computeDogYears(myAge)
	fmt.Println("My age in dog years:",myDogsAge)

	totalPriceFr := vatTaxFr(500)
	fmt.Println("The price of a gaming console in France with the TVA added:",totalPriceFr)
}

func computeDogYears(humanYears int) int {
	dogYears := humanYears * 7
	return dogYears
}

func vatTaxFr(price float32) float32 {
	vatFr := (price / 200) * 100 * .20
	return price + vatFr
}

func salesTaxPhila(salesTax, price float32) float32 {
	salesTax = price * (salesTax / 100)
	return price + salesTax
}
