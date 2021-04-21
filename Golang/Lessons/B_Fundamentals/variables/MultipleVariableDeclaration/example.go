package main

import "fmt"

func main() {
	var part1, partB string
	part1 = "Resident Evil"
	partB = "2"

	// Above, we declared both part1 and part2 on the same line both with the same type
	// Using this syntax, both variables must be the same type (in this case they are "string")

	fmt.Println(part1, partB)

	// if we already know what values we want to assign to our variable,
	// We can use the :=

	quote, reality := "is sales tax or TVA in France is 20% ?", true
	fmt.Print(quote, ":\t ", reality)

}

// Go always use to declare multiple variables on a single line
