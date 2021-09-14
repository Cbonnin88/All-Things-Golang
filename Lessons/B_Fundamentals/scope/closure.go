package main

import "fmt"

func main() {
	// num1 is in our outer scope enclosing our inner scope
	num1 := 45
	fmt.Println(num1)
	// Here we have our inner scope, we can access num1 in the inner scope
	{
		fmt.Println(num1)
		num2 := 22
		fmt.Println(num2)
	}
	// num2 cannot be accessed in the outer scope
}
