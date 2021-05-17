package main

import "fmt"

func main() {
	// create a slice:
	var num []int = []int{5, 10, 15, 20, 25, 30, 35, 40}
	// using the range keyword
	fmt.Println("Using the range keyword:")
	for i, element := range num {
		// element = num[i]
		fmt.Printf("%d: %d\n", i, element)
	}

	fmt.Println("Finding duplicate using the range method:")
	var doubleNum []int = []int{5, 10, 15, 20, 25, 10, 25, 30}
	for i, element := range doubleNum {
		for j, element2 := range doubleNum {
			// Creating the condition to detect the duplicate numbers:
			if element == element2 && j > i {
				// Prints out the duplicate numbers 10 and 25 in our slice
				fmt.Println("\nThe duplicate number in our slice is: ", element)
				fmt.Println(doubleNum, "\n")
				// prints out our slice
			}
		}
	}
	// Simplified method for finding the duplicate element
	fmt.Println("Using a simplified method to find the duplicate element:")
	var countries []string = []string{"France", "Germany", "Finland", "France", "United Kingdom"}
	for i, country := range countries {
		for j := i + 1; j < len(countries); j++ {
			country2 := countries[j]
			if country2 == country {
				fmt.Println(countries)
				fmt.Println("The duplicate element in our string slice is : ", country)
			}
		}
	}
}
