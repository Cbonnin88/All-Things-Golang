package main

import "fmt"

// Loops in Go :
func main() {

	// A simple Loop:
	fmt.Println("Simple Loop from zero to five:")
	for x := 0; x <= 5; x++ {
		fmt.Println(x)
	}

	// For each or range Loop:
	fmt.Println("Range Loops")
	rangeVar := []int{10, 20, 30, 40}
	for i, j := range rangeVar {
		fmt.Println(i, j)
	}
}
