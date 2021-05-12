package main

import "fmt"

func main() { //nolint:typecheck
	var arr [5]int
	arr[0] = 10
	arr[1] = 20
	arr[2] = 30
	arr[3] = 40
	arr[4] = 50

	fmt.Println("a basic array: ", arr)

	// Shorthand for entering values in an array:
	var arr2 = [4]string{"House", "tree", "ghost", "broom"}
	fmt.Println("The shorthand version: ", arr2)

	// This lets the compiler guess the index number:
	var arr3 = [...]int{11, 55, 32, 100, 150}
	fmt.Println("The compiler guesses the index number: ", arr3)

	// Using Loops to access elements in an array:
	var arr4 = [5]string{"House", "kids", "spouse", "car", "love"}

	fmt.Println("Using loops with an array:")
	for i := 0; i < 5; i++ {
		fmt.Println(arr4[i])
	}

	// using len() method, you can find the length of an array :
	fmt.Println("The length of array 1 is: ", len(arr))
	fmt.Println("The length of array 2 is: ", len(arr2))
	fmt.Println("The length of array 3 is: ", len(arr3))
	fmt.Println("The length of array 4 is: ", len(arr4))
}
