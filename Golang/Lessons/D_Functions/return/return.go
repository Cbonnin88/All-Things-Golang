package main

import "fmt"

func main() {
	// Calling our newCar() function in the main
	priceTag := newCar()
	fmt.Print(priceTag)
}

func newCar() string {
	var choice string
	car1 := 25000
	car2 := 50000
	car3 := 75000
	fmt.Print("Enter you budget amount: ")
	var budget int
	_, err := fmt.Scan(&budget)
	if err != nil {
	}

	switch {
	case car1 <= budget && budget < car2:
		choice = "The Lexus is in your price range"
	case car2 <= budget && budget < car3:
		choice = "The Ferrari is in your price range"
	case car3 <= budget && budget > car2:
		choice = "The Audi is in your price range"
	default:
		choice = "These models are too expensive"
	}
	return choice // our return statement
}
