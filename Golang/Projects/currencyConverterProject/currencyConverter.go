package main

import "fmt"

func main() {
	for {
		fmt.Println("Which currency would you like to convert ?:")
		fmt.Println("\t1:USD \t2:EURO \t3:GBP \t4:JPY ")
		var choice1 int
		_, _ = fmt.Scan(&choice1)
		fmt.Println("Into what currency ?:")
		fmt.Println("\t5:USD \t6:EURO \t7:GBP \t8:JPY ")
		var choice2 int
		_, _ = fmt.Scan(&choice2)
		fmt.Print("How much money do you want to convert ?: ")
		var amount float32
		_, _ = fmt.Scan(&amount)

		var dollarToEuro = usdToEuro(amount)
		var dollarToGbp = usdToGbp(amount)
		var dollarToJpy = usdToJpy(amount)
		var dollarToDollar = usdToUsd(amount)

		var euroToDollar = euroToUSD(amount)
		var euroToGBP = euroToGBP(amount)
		var euroToJPY = euroToJPY(amount)
		var euroToEuro = euroToEuro(amount)

		var gbpToDollar = gbpToUSD(amount)
		var gbpToEuro = gbpToEuro(amount)
		var gbpToJPY = gbpToJPY(amount)
		var gbpToGbp = gbpToGbp(amount)

		var jpyToDollar = jpyToUsd(amount)
		var jpyToEuro = jpyToEuro(amount)
		var jpyToGbp = jpyToGbp(amount)
		var jpyToJpy = jpyToJpy(amount)

		switch {
		case choice1 == 1 && choice2 == 6:
			fmt.Printf("\n%.2f USD = %.2f EUR", amount, dollarToEuro)
		case choice1 == 1 && choice2 == 7:
			fmt.Printf("\n%.2f USD = %.2f GBP", amount, dollarToGbp)
		case choice1 == 1 && choice2 == 8:
			fmt.Printf("\n%.2f USD = %.2f JPY", amount, dollarToJpy)
		case choice1 == 1 && choice2 == 5:
			fmt.Printf("\n%.2f USD = %.2f USD", amount, dollarToDollar)
		}

		switch {
		case choice1 == 2 && choice2 == 5:
			fmt.Printf("\n%.2F EUR = %.2f USD", amount, euroToDollar)
		case choice1 == 2 && choice2 == 7:
			fmt.Printf("\n%.2f EUR = %.2f GBP", amount, euroToGBP)
		case choice1 == 2 && choice2 == 8:
			fmt.Printf("\n%.2f EUR == %.2f JPY", amount, euroToJPY)
		case choice1 == 2 && choice2 == 6:
			fmt.Printf("\n%.2f EUR = %.2f EUR", amount, euroToEuro)
		}

		switch {
		case choice1 == 3 && choice2 == 5:
			fmt.Printf("\n%.2f GBP = %.2f USD", amount, gbpToDollar)
		case choice1 == 3 && choice2 == 6:
			fmt.Printf("\n%.2f GBP = %.2f EUR", amount, gbpToEuro)
		case choice1 == 3 && choice2 == 8:
			fmt.Printf("\n%.2f GBP = %.2f JPY", amount, gbpToJPY)
		case choice1 == 3 && choice2 == 7:
			fmt.Printf("\n%.2f GBP = %.2f GBP", amount, gbpToGbp)
		}

		switch {
		case choice1 == 4 && choice2 == 5:
			fmt.Printf("\n%.2f JPY = %.2f USD", amount, jpyToDollar)
		case choice1 == 4 && choice2 == 6:
			fmt.Printf("\n%.2f JPY = %.2f EUR", amount, jpyToEuro)
		case choice1 == 4 && choice2 == 7:
			fmt.Printf("\n%.2f JPY = %.2f GBP", amount, jpyToGbp)
		case choice1 == 4 && choice2 == 8:
			fmt.Printf("\n%.2f JPY = %.2f JPY", amount, jpyToJpy)
		}
		fmt.Println("\n\nWould you like to convert another currency ? (Y/N)")
		var answer string
		_, _ = fmt.Scan(&answer)
		if answer == "N" {
			break
		} else if answer == "Y" {
			continue
		}
	}
}

func usdToEuro(convertUsd float32) float32 {
	const exchangeRateEuro float32 = .82
	var rateEURO = convertUsd * exchangeRateEuro
	return rateEURO
}
func usdToGbp(convertUsd2 float32) float32 {
	const exchangeRateGBP float32 = .71
	var rateGBP = convertUsd2 * exchangeRateGBP
	return rateGBP
}
func usdToJpy(convertUsd3 float32) float32 {
	const exchangeRateJPY float32 = 105.42
	var rateJPY = convertUsd3 * exchangeRateJPY
	return rateJPY
}
func usdToUsd(noConvert float32) float32 {
	const sameRate float32 = 1
	var noRate = noConvert * sameRate
	return noRate
}

func euroToUSD(euroConvert float32) float32 {
	const exchangeRateUSD float32 = 1.21
	var rateUSD = euroConvert * exchangeRateUSD
	return rateUSD
}

func euroToGBP(euroConvert2 float32) float32 {
	const exchangeRateGBP float32 = .86
	var rateEuroToGbp = euroConvert2 * exchangeRateGBP
	return rateEuroToGbp
}

func euroToJPY(euroConvert3 float32) float32 {
	const exchangeRateJPY float32 = 127.71
	var rateEuroToJpy = euroConvert3 * exchangeRateJPY
	return rateEuroToJpy
}

func euroToEuro(sameCurrency float32) float32 {
	const noChange float32 = 1
	var euroToEuro = sameCurrency * noChange
	return euroToEuro
}

func gbpToUSD(gbpConvert float32) float32 {
	const exchangeRateUSD float32 = 1.40
	var rateGbpToUsd = gbpConvert * exchangeRateUSD
	return rateGbpToUsd
}

func gbpToEuro(gbpConvert2 float32) float32 {
	const exchangeRateEuro float32 = 1.40
	var rateGbpToEuro = gbpConvert2 * exchangeRateEuro
	return rateGbpToEuro
}

func gbpToJPY(gbpConvert3 float32) float32 {
	const exchangeRateJPY float32 = 147.67
	var rateGbpToJPY = gbpConvert3 * exchangeRateJPY
	return rateGbpToJPY
}

func gbpToGbp(noChange float32) float32 {
	const noChange1 float32 = 1
	var sameMoney = noChange * noChange1
	return sameMoney
}

func jpyToUsd(jpyConvert float32) float32 {
	const exchangeRateUSD float32 = 0.009
	var rateJpyToUsd = jpyConvert * exchangeRateUSD
	return rateJpyToUsd
}

func jpyToEuro(jpyConvert2 float32) float32 {
	const exchangeRateEuro float32 = 0.007
	var rateJpyToEuro = jpyConvert2 * exchangeRateEuro
	return rateJpyToEuro
}

func jpyToGbp(jpyConvert3 float32) float32 {
	const exchangeRateGBP float32 = 0.006
	var rateJpyToGBP = jpyConvert3 * exchangeRateGBP
	return rateJpyToGBP
}

func jpyToJpy(noChangeJpy float32) float32 {
	const exchangeJPY float32 = 1
	var jpyToJpy = noChangeJpy * exchangeJPY
	return jpyToJpy
}
