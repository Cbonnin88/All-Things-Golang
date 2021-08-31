package main

import "fmt"

func main() {
	result := Math(3)
	fmt.Println(result)
}

func Math(x int) (result int) {
	result = x * 2
	return result
}
