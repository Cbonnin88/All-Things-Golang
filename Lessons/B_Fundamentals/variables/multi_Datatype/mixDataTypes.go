package main

import "fmt"

func main() {
	var cityName = "Chartres"
	houseCost := 200000
	var percentRate float32 = 1.6
	fmt.Printf("We bought a house in "+cityName+" last month. We only paid %d euros for a 4 bedroom home. The mortgage rate was only %.2f percent !!!\n", houseCost, percentRate)
	fmt.Println("The length of the word 'Chartres' is: ", len(cityName))

}
