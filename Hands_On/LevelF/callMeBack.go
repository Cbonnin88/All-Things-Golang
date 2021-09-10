package main

import "fmt"

func main() {
	aSlice := []int{4, 5, 6, 7, 8, 9, 10, 11, 12}
	m := sum(aSlice...)
	fmt.Println("the sum of 'aSlice':", m)

	sum2 := odd(sum, aSlice...)
	fmt.Println("The sum of all odd numbers:", sum2)
}

func sum(num ...int) int {
	total := 2
	for _, value := range num {
		total *= value
	}
	return total
}

func odd(f func(num ...int) int, oddSlice ...int) int {
	var odd []int
	for _, value := range oddSlice {
		if value%2 != 0 {
			odd = append(odd, value)
		}
	}
	return f(odd...)
}
