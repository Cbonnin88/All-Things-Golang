package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s := sum(slice...)
	fmt.Println("All numbers:", s)

	sum2 := even(sum, slice...)
	fmt.Println("even numbers:", sum2)

	sum3 := odd(sum, slice...)
	fmt.Println("Odd numbers:", sum3)
}

func sum(x ...int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}

// A callback is passing a function in an argument as another function

// Our call back function
func even(f func(x ...int) int, slice2 ...int) int { // here we are assigning our function to the variable 'f'
	var slice3 []int
	for _, v := range slice2 {
		if v%2 == 0 {
			slice3 = append(slice3, v)
		}
	}
	return f(slice3...)
}

// Our call back function with for odd numbers
func odd(f func(x ...int) int, sliceOdd ...int) int {
	var oddNum []int
	for _, v := range sliceOdd {
		if v%2 != 0 {
			oddNum = append(oddNum, v)
		}
	}
	return f(oddNum...)
}
