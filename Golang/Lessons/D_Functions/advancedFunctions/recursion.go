package main

import "fmt"

func main() { //nolint:typecheck
	// Recusion is when a function calls itself a certain number of times and then stops
	fmt.Println("The factoral of five is:",5 * 4 * 3 * 2 * 1)
	fact := factorial(6)
	fmt.Println("The factorial of six:",fact)

	fact2 := loopFact(7)
	fmt.Println("The factorial of seven, using a loop:",fact2)

}
// Using recusion to calculate our factorial:
func factorial(num int) int {
	if num == 0 {
		return 1
	}
	return num * factorial(num - 1)
}
// Using a loop
func loopFact(num int) int {
		total := 1
		for ; num > 0; num-- {
			total *= num
		}
		return total
}
