package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Print("enter a modulus: ")
	var digit int
	_, err := fmt.Scan(&digit)
	if err != nil {
		log.Print(err)
	}
	fmt.Print("enter a number: ")
	var num int
	_, err = fmt.Scan(&num)
	if err != nil {
		log.Print("Please enter a valid number")
	}
	for n := num; n <= 100; n++ {
		fmt.Printf("When %v is divided by %v, the remainder is %v\n", n, digit, n%digit)
	}
}
