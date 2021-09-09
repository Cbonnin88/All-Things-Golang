package main

import (
	"fmt"
	"log"
)

// Num1 and Num2 hold our int values
var num1 int
var num2 int

func main() {
	fmt.Print("Enter a number: ")
	// We use user input for our int value
	_, err := fmt.Scan(&num1)
	if err != nil {
		log.Println("Please enter a number !!!")
	}

	fmt.Print("Enter a second number: ")
	_, err = fmt.Scan(&num2)
	if err != nil {
		log.Println("Please enter a number !!!")
	}

	fmt.Println(num1, "*", num2, "=", multiply(num1, num2))
}

// Multiply multiplies an unlimited value of type int values
func multiply(num ...int) int {
	var ans int
	for _, num1 := range num {
		ans = num1 * num2
	}
	return ans
}
