package main

import "fmt"

func taxCalculatorEU(price, vatTax float32) float32 {
	vatEu := (price / 200) * 100 * vatTax

	return vatEu + price
}

func main() {

	var a, b, c, d float32
	a = .20
	b = .17
	c = .24
	d = .16

	var france float32 = taxCalculatorEU(500, a)
	var luxembourg float32 = taxCalculatorEU(500, b)
	var finland float32 = taxCalculatorEU(500, c)
	var germany float32 = taxCalculatorEU(500, d)

	fmt.Println("The new PlayStation 5 is 5OO EUR without VAT inclued, let's see how much it is in four different European Countries when we include the VAT tax:")
	fmt.Printf("France (20 percent VAT): %.2f\n", france)
	fmt.Printf("Luxembourg (17 percent VAT): %.2f\n", luxembourg)
	fmt.Printf("Finland (24 percent VAT): %.2f\n", finland)
	fmt.Printf("Germany (16 percent VAT): %.2f\n", germany)
}
