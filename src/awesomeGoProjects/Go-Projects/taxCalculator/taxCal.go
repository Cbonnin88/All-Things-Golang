package main

import "fmt"

func main() {
	france, tva, philly, salesTax := "France has a VAT or sales tax of", 20, " while Philadelphia has a sales tax of", 8
	fmt.Print(france, " ", tva, "%", philly, " ", salesTax, "%\n")
	fmt.Println("Here's the price difference between the French VAT vs the Philadelphia sales tax:\n ")

	fmt.Print("Enter an amount: ")
	var priceOfGame float32
	_, _ = fmt.Scan(&priceOfGame)
	var finalPrice = getPriceWithTax(priceOfGame)
	var totalPricePh = getPricePhillyTax(priceOfGame)
	var priceDifference = getDifferencePrice(finalPrice, totalPricePh)
	var dollarDifference = euroToDollars(priceDifference)

	fmt.Printf("The price of this item, with the French VAT included,is %.2f euros\n", finalPrice)
	fmt.Printf("The price of this item, with Philadelphia sales tax included, is %.2f dollars\n", totalPricePh)
	fmt.Printf("The price difference between these two items is %.2f euros or %.2f US dollars", priceDifference, dollarDifference)

}

func getPriceWithTax(price float32) float32 {
	const vat float32 = .20

	var totalPrice = (price/100.2)*100*vat + price

	return totalPrice
}

func getPricePhillyTax(price float32) float32 {
	const phSalesTax float32 = 0.08
	var totalPricePh = price*phSalesTax + price
	return totalPricePh
}
func getDifferencePrice(price float32, price2 float32) float32 {
	var differencePrice = price - price2
	return differencePrice
}
func euroToDollars(price float32) float32 {
	const exchangeRate float32 = 1.18
	var rate = price * exchangeRate
	return rate
}
