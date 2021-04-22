package main

import "fmt"

func main() {
	firstFunction()
	secondFunction(100)
	add(10, 20)

	answer := multiply(100, 200)
	fmt.Println(answer)

	ans1, ans2 := math(100, 1200)
	fmt.Println(ans1, ans2)

	ans3, ans4 := math2(12, 24)
	fmt.Println(ans3, ans4)

	def1, def2 := deferring(10, 30)
	fmt.Println(def1, def2)
}

// a simple function
func firstFunction() {
	fmt.Println("This is my first function in Golang")
}

// a simple function with a parameter
func secondFunction(a int) {
	fmt.Println(a)
}

// a simple function with several parameters
func add(b int, c int) {
	fmt.Println(b + c)
}

// return a value
func multiply(x int, y int) int {
	fmt.Println(x * y)
	return x * y
}

// returning multiple values
func math(num1 int, num2 int) (int, int) {
	return num1 * num2, num2 / num1
}

// Labeling return variables
func math2(c int, d int) (z1, z2 int) {
	z1 = c + d
	z2 = d * c
	return
}

// Using a defer statement
func deferring(x, y int) (a1, a2 int) {
	defer fmt.Println("This is a deferred statement")
	a1 = x + y
	a2 = x - y
	fmt.Println("This is NOT a deferred statement")
	return
}
