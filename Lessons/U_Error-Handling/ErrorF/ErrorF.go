package main

import (
	"fmt"
)

var num int

func main() {

	fmt.Print("Enter a number: ")
	_, err := fmt.Scan(&num)
	if err != nil {
		panic(err)
	}

	_, err = numbers(num)
	if err != nil {
		panic(err)
	}
	fmt.Printf("This is your number: %d", num)
}

func numbers(posNum int) (int, error) {
	if posNum < 0 {
		numErr := fmt.Errorf("error: %v is not a positive number", posNum) // Here we create our custom errors with detailed information
		return num, numErr
	}
	return num, nil
}
