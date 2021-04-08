package main

import "fmt"

func taxCalculatorEU(price, vatTax float32) float32 {
	vatEu := (price / 200) * 100 * vatTax

	return vatEu + price
}

func euroToUsd(price, convertRate float32) float32 {
	conversion := price * convertRate

	return conversion
}

func pricePh() float32 {
	salesTax := euroToUsd(500, .83) * (.08 / 100)
	var totalPricePh float32 = salesTax + euroToUsd(500, .83)
	return totalPricePh
}

func goingBack() float32 {
	var rate float32 = 1.21
	return rate
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

	var differenceFr, differenceLu, differenceFI, differenceGe float32
	differenceFr = taxCalculatorEU(500, a) - pricePh()
	differenceLu = taxCalculatorEU(500, b) - pricePh()
	differenceFI = taxCalculatorEU(500, c) - pricePh()
	differenceGe = taxCalculatorEU(500, d) - pricePh()

	var dollarFr, dollarLu, dollarFI, dollarGe float32
	dollarFr = differenceFI * goingBack()
	dollarLu = differenceLu * goingBack()
	dollarFI = differenceFI * goingBack()
	dollarGe = differenceGe * goingBack()

	fmt.Println("The new PlayStation 5 is €5OO without VAT inclued, let's see how much it costs in four different European Countries when we include the VAT tax:")
	fmt.Printf("France (20 percent VAT):   €%.2f\n", france)
	fmt.Printf("Luxembourg (17 percent VAT): €%.2f\n", luxembourg)
	fmt.Printf("Finland (24 percent VAT): €%.2f\n", finland)
	fmt.Printf("Germany (16 percent VAT): €%.2f\n", germany)
	fmt.Printf("The Playstation 5 in the United States costs $500 (Philadelphia Sales tax included), in Euros this amount would be €%.2f\n", pricePh())
	fmt.Println("This is the difference, in Euros, between the USA and each country:")
	fmt.Printf("France: €%.2f or $%.2f\n", differenceFr, dollarFr)
	fmt.Printf("Luxembourg: €%.2f or $%.2f\n", differenceLu, dollarLu)
	fmt.Printf("Finland: €%.2f or $%.2f\n", differenceFI, dollarFI)
	fmt.Printf("Germany: €%.2f or $%.2f\n", differenceGe, dollarGe)

}
