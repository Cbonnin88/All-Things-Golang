package main

import "fmt"

func main() {
	slice := []int{4, 8, 12, 16, 20, 24, 28, 32}
	add(slice...)
}

// our sum function, using the range method
func add(sum ...int) int {
	fmt.Println(sum)

	add := 0
	for i, v := range sum {
		add += v
		fmt.Println("for the item in index position", i, "we are adding", v, "to the total, which is now", add)
	}
	fmt.Println("The total is:", add)
	return add
}
