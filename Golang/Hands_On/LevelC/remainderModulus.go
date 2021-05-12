package main

import "fmt"

func main() { //nolint:typecheck
	fmt.Print("enter a modulus: ")
	var digit int
	_, err := fmt.Scan(&digit)
	if err != nil {
		return
	}
	fmt.Print("enter a number: ")
	var num int
	_, err = fmt.Scan(&num)
	if err != nil {
		return
	}
	for n := num; n <= 100; n++ {
		fmt.Printf("When %v is divided by %v, the remainder is %v\n", n, digit, n%digit)
	}
}
