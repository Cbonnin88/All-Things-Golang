package main

import "fmt"

func main() {
	daysOnVacation := 15   // daysOnVacation becomes an integer because 15 is an integer
	var monthsInAYear = 12 // same rule for monthsInAYear
	fmt.Println("You have spent", daysOnVacation*monthsInAYear, "weeks on vacation gurl")

	// Example 2:
	title := "Golang Guide"
	var author = "Golang Guru Lady"
	fmt.Println("The title of this book is", title, "by best selling author", author)
}
