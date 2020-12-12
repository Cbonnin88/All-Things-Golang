package main

import "fmt"

func newCar() string {
	var choice string
	car1 := 25000
	car2 := 50000
	car3 := 75000
	budget := 55000

	if car1 <= budget && budget < car2 {
		choice = "The Lexus is in your price range"
	} else if car2 <= budget && budget < car3 {
		choice = "The Ferrari is in your price range"
	} else if car3 <= budget && budget > car2 {
		choice = "The Audi is in your price range"
	} else {
		choice = "These cars are too expensive"
	}

	return choice
}

func main() {

	var priceTag string
	priceTag = newCar()
	fmt.Printf(priceTag)
}
