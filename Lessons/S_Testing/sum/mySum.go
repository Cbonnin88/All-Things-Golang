package main

import "fmt"

func main() {
	fmt.Println("1 + 1 =", sum(1, 1))
	fmt.Println("100 + 200 =", sum(100, 200))
	fmt.Println("20192 + 3324 =", sum(20192, 3324))
	fmt.Println("6678 + 234 =", sum(6678, 234))

}

func sum(num ...int) int {
	mySum := 0
	for _, v := range num {
		mySum += v
	}
	return mySum
}
