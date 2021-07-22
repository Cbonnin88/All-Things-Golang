package main

import "fmt"

func main() { //nolint:typecheck
	// We can loop over a slice just like an array:

	// for loop with a slice using "Range":
	slice := []int{10, 20, 30, 40, 50}
	for i, v := range slice {
		fmt.Printf("index: %v\nvalue: %v\n", i, v)

	}

}
