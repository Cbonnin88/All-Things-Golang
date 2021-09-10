package main

import "fmt"

func main() {
	// Everything in Go is passed by VALUE
	firstFunction()
	secondFunction(100)
	add(10, 20)
	unlimited(2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12)
	multiply(100, 200)

	ans1, ans2 := math(100, 1200)
	fmt.Println("Returning multiple values:", ans1, ans2)

	ans3, ans4 := math2(12, 24)
	fmt.Println("Labeling return values:", ans3, ans4)

	def1, def2 := deferring(10, 30)
	fmt.Println(def1, def2)
}

// a simple function
func firstFunction() {
	fmt.Println("This is my first function in Golang")
}

// a simple function with a parameter
func secondFunction(a int) {
	fmt.Println("this is a function with a parameter:", a)
}

// a simple function with several parameters
func add(b int, c int) {
	fmt.Println("Multiple parameters that do addition:", b+c)
}

// a simple function with unlimited parameters
func unlimited(num ...int) {
	fmt.Println("with unlimited parameters:", num)
}

// return a value
func multiply(x int, y int) int {
	fmt.Println("Returning a value:", x*y)
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

// Defer statement
func deferring(x, y int) (a1, a2 int) {
	defer fmt.Println("This is a deferred statement")
	a1 = x + y
	a2 = x - y
	fmt.Println("This is NOT a deferred statement")
	return
}
