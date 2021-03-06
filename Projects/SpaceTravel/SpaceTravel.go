package main

import "fmt"

func fuelGauge(fuel int) {
	fmt.Printf("We have %v percent of fuel left\n", fuel)
}

func calculateFuel(planet string) int {
	var fuel int

	switch planet {
	case "Venus":
		fuel = 300000
	case "Mercury":
		fuel = 500000
	case "Mars":
		fuel = 700000
	default:
	}
	return fuel //nolint:wsl
}
func greetPlanet(planet string) {
	fmt.Println("Welcome to planet", planet, "traveler !!!")
}

func cantFly() {
	fmt.Println("We do not have the available amount of fuel to fly there.") //nolint:forbidigo
}

func flyToPlanet(planet string, fuel int) int {
	var fuelRemaining, fuelCost int
	fuelRemaining = fuel
	fuelCost = calculateFuel(planet)

	if fuelRemaining >= fuelCost {
		greetPlanet(planet)
		fuelRemaining -= fuelCost
	} else {
		cantFly()
	}
	return fuelRemaining
}

func main() {
	var fuel int
	fuel = 1000000
	planetChoice := "Venus"
	fuel = flyToPlanet(planetChoice, fuel)
	fuelGauge(fuel)

}
