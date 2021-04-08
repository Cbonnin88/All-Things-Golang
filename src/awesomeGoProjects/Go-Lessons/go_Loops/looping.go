package main

import "fmt"

// Loops in Go :
func main() {
	sum := 0

	for i := 1; i < 20; i++ {
		sum += i
	}

	// Another way of writing a for loop in Go:
	for x := 0; x <= 5; x++ {
		fmt.Println(x)
	}

	fmt.Println(sum)
	fmt.Println(loops())
}

func loops() int {
	fmt.Print("Enter a number: ")
	var num int
	_, _ = fmt.Scan(&num)
	for i := 5; i < 10; i++ {
		num += i
	}
	return num
}
