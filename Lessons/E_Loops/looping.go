package main

import (
	"fmt"
	"time"
)

// Loops in Go :
func main() { //nolint:typecheck

	// A simple Loop:
	fmt.Println("Simple incrementing Loop from zero to five:")
	for x := 0; x <= 5; x++ {
		fmt.Println(x)
	}
	fmt.Println("")

	// For each or range Loop:
	fmt.Println("Range Loops")
	rangeVar := []int{10, 20, 30, 40}
	for i, j := range rangeVar {
		fmt.Println(i, j)
	}
	fmt.Println("")

	// Counting down with a While Loop:
	fmt.Println("Lets begin the count down shall we?")
	for i := 10; i > 0; i-- {
		fmt.Println(i)
		time.Sleep(time.Second)
	}
	fmt.Println("Lift Off")
}
