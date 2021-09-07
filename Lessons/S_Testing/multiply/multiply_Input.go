package main

import (
	"fmt"
	"log"
)

var num1 int
var num2 int

func main() {
	fmt.Print("Enter a number: ")
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

func multiply(num ...int) int {
	var ans int
	for _, num1 := range num {
		ans = num1 * num2
	}
	return ans
}
