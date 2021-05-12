package main

import "fmt"

func main() {
	fmt.Print("Please enter your desired yearly salary: ")
	var salary int
	_, err := fmt.Scan(&salary)
	if err != nil {
		return
	}
	switch salary {
	case 45000:
		fmt.Println("Here's our offer of 45K")
	case 55000:
		fmt.Println("Here's our offer of 55K")
	case 65000:
		fmt.Println("Here's our offer of 65K")
	case 70000:
		fmt.Println("Here's our offer of 70K")
	default:
		fmt.Println("We will get back to you with an offer in two days")
	}
	fmt.Println("")

	// An example of a switch statement using a string:
	fmt.Print("Please type in your desired article:")
	var clothing string
	_, err = fmt.Scan(&clothing)
	if err != nil {
		return
	}

	switch clothing {
	case "shirt":
		fmt.Println("Here's our collection of shirts")
	case "shoes":
		fmt.Println("Here's our shoe collection")
	case "jeans":
		fmt.Println("Here's our latest collection of jeans")
	case "dress":
		fmt.Println("Here's our summer dresses")
	case "jacket":
		fmt.Println("Here's our fall line up of jackets")
	default:
		fmt.Println("Sorry, we don't have that currently in stock")
	}

	fmt.Println()
	// Using a for loop in a Switch Statement with fallthrough:
	for day := 1; day <= 12; day++ {
		fmt.Println("On the", day, "day of Christmas, my true love sent to me:")
		switch day {
		case 12:
			fmt.Println("Twelve Drummers drumming,")
			fallthrough
		case 11:
			fmt.Println("Eleven pipers piping,")
			fallthrough
		case 10:
			fmt.Println("Ten lords a-leaping,")
			fallthrough
		case 9:
			fmt.Println("Nine ladies dancing,")
			fallthrough
		case 8:
			fmt.Println("Eight maids a-milking,")
			fallthrough
		case 7:
			fmt.Println("Seven swans a-swimming,")
			fallthrough
		case 6:
			fmt.Println("Six geese a-laying,")
			fallthrough
		case 5:
			fmt.Println("Five golden rings,")
			fallthrough
		case 4:
			fmt.Println("Four calling birds,")
			fallthrough
		case 3:
			fmt.Println("Three french hens,")
			fallthrough
		case 2:
			fmt.Println("Two turtle doves, and,")
			fallthrough
		case 1:
			fmt.Println("A partridge in a pear tree")

			// A fallthrough statement transfers control to the next case
			// It may only be used as the final statement in a clause
		}
	}
}
