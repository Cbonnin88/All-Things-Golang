package main

import "fmt"

func main() {
	test()
	test2("Running will get you no where!!!")
	test3(100, 200)
	test4(1.1, 2.2, 3.3, 4.4, 5.5, 6.6, 7.7, 8.8, 9.9, 10.1)
	fmt.Println(test5(100, 10))
	fmt.Println(test6(10.4, 20.22))
	fmt.Println(test7(10, 20))
}

func test() {
	fmt.Println("A simple Function")
}

func test2(a string) {
	fmt.Println("Alcina says:", a)
}

func test3(b, c int) {
	fmt.Println(b + c)
}

func test4(num ...float32) {
	fmt.Println(num)
}

func test5(a, b int) int {
	return a / b
}

func test6(e, f float32) (float32, float32) {
	return e * f, f / e
}

func test7(h, i int) (j, k int) {
	defer fmt.Println("I am a defer statement")
	j = h + i
	k = i * h
	fmt.Println("I am NOT a defer statement")
	return j, k
}
