package main

import "fmt"

func main() {
	var bookPrice1 = bookStore01(8.79)

	var publisher, writer, artist, title string
	title = "Mr.GoToSleep"
	writer = "Tracey Hatchet"
	artist = "Jewel Tampson"
	publisher = "DizzyBooks Publishing Inc."
	var year, pageNumber uint
	year = 1997
	pageNumber = 14
	var grade float32
	grade = 6.5

	fmt.Printf("%v was written by %v and drawn by %v.The publishing house is %v,.This book was written in %d and it has %d pages. The book gets a %.f because it is a used copy. The price for this book is %.2f euros, VAT included\n\n", title, writer, artist, publisher, year, pageNumber, grade, bookPrice1)

	title = "Epic Vol 1"
	writer = "Ryan N. Shawn"
	artist = "Phoebe Paperclips"
	publisher = "DizzyBooks Publishing Inc."
	year = 2013
	pageNumber = 160
	grade = 9.0

	var bookPrice2 = bookStore02(21.98)

	fmt.Printf("%v was written by %v and drawn by %v.The publishing house is %v,.This book was written in %d and it has %d pages. The book gets a %.f because it is a used copy. The price for this book is %.2f euros, VAT included\n\n", title, writer, artist, publisher, year, pageNumber, grade, bookPrice2)

	title = "Fade Away"
	writer = "Derrick Allison"
	artist = "Mayson Grey"
	publisher = "RockHouse Publishing Company"
	year = 2020
	pageNumber = 776
	grade = 9.0

	var bookPrice3 = bookStore03(45.56)

	fmt.Printf("%v was written by %v and drawn by %v.The publishing house is %v,.This book was written in %d and it has %d pages. The book gets a %.f because it is a used copy. The price for this book is %.2f euros, VAT included\n\n", title, writer, artist, publisher, year, pageNumber, grade, bookPrice3)
}
func bookStore01(price float32) float32 {
	const price1 float32 = 8.79
	totalPrice := (price/100.2)*100*.20 + price1
	return totalPrice
}

func bookStore02(price float32) float32 {
	const price2 = 21.98
	totalPrice2 := (price/100.2)*100*.20 + price2
	return totalPrice2
}

func bookStore03(price float32) float32 {
	const price3 = 45.56
	totalPrice3 := (price/100.2)*100*.20 + price3
	return totalPrice3
}
