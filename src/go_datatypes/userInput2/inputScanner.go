package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	// the %q puts the string into "" marks
	// this is another way to get userInput via the console in Go

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter your birth year: ")
	scanner.Scan()
	input, _ := strconv.ParseInt(scanner.Text(), 10, 64) // to convert an integer into a string while using the input console
	fmt.Printf("You will be %d years old at the end of 2021", 2021-input)
}
