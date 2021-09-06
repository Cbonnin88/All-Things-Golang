package main

import (
	"fmt"
	"log"
)

var min int // Here I declare my 'cost' variable
var max int

func main() {
	fmt.Print("Enter minimum price: ") // Declaring the User input
	_, err := fmt.Scan(&min)
	if err != nil {
		log.Fatalln(err)
	}

	_, err = newHouseMin(min) // Here I check the error for the newHouse function
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Print("Enter maximum price: ")
	_, err = fmt.Scan(&max)
	if err != nil {
		log.Fatalln(err)
	}
	_, err = newHouseMax(max)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("here are some houses in your area according to your price range of %v and %v\n", min, max)

}

func newHouseMin(price int) (int, error) {
	if price < 0 {
		priceErr := fmt.Errorf("error: %v is a negative number", price) // I create a custom error message
		return min, priceErr                                            // if less than zero, it returns the min, and the error
	}
	return min, nil // here we return the cost
}

func newHouseMax(price2 int) (int, error) {
	if price2 < 0 {
		priceErr2 := fmt.Errorf("error: %v is a negative number", price2)
		return max, priceErr2
	}
	return max, nil
}
