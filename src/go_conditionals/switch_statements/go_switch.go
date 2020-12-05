package main

import "fmt"

func main() {
	salary := "35K"

	switch salary {
	case "45K" :
		fmt.Println("Here's our offer of 45K")
	case "55K":
		fmt.Println("Here's our offer of 55K")
	case "65K":
		fmt.Println("Here's our offer of 65K")
	case "70K":
		fmt.Println("Here's our offer of 70K")
	default:
		fmt.Println("We will get back to you with an offer in two days")
	}
}
