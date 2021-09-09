package main

import "fmt"

func main() {
	result := math(3)
	fmt.Println(result)
}

// Math multiplies an unlimited number of int values by two
func math(x int) (result int) {
	result = x * 2
	return result
}
