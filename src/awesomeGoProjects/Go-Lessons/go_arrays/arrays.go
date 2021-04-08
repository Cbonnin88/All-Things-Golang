package main

import "fmt"

func main() {
	var arr [5]int
	arr[0] = 10
	arr[1] = 20
	arr[2] = 30
	arr[3] = 40
	arr[4] = 50

	// Shorthand for entering values in an array:
	var arr2 = [4]string{"House", "tree", "ghost", "broom"}

	// This lets the compiler guess the index number:
	var arr3 = [...]int{11, 55, 32, 100, 150}

	// Using Loops to access elements in an array:
	var arr4 = [5]string{"House", "kids", "spouse", "car", "love"}

	for i := 0; i < 5; i++ {
		fmt.Println(arr4[i])
	}

	fmt.Println(arr)
	fmt.Println(arr2)
	fmt.Println(arr3)

	// using len() method, you can find the length of an array :
	fmt.Println("The length of array 1 is: ", len(arr))
	fmt.Println("The length of array 2 is: ", len(arr2))
	fmt.Println("The length of array 3 is: ", len(arr3))
	fmt.Println("The length of array 4 is: ", len(arr4))
}
