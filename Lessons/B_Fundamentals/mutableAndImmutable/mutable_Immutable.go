package main

import "fmt"

func main() {
	// Example of a mutable datatype:
	var mutable = 30
	y := mutable
	y = 10
	fmt.Println(mutable, y)
	// Mutable datatype as a slice:
	var slice = []float32{19.7, 12.8, 29.4}
	y1 := slice
	y1[0] = 124.5
	fmt.Println(slice, y1)
	// Maps are also mutable datatype:
	var x = map[string]int{"Hello There": 4}
	y3 := x
	x["7"] = 7
	y3["y"] = 50
	fmt.Println(x, y3)
	// Using a mutable array:
	var arr = [3]string{"cow", "egg", "french"}
	yArr := arr
	yArr[0] = "Zombie"
	fmt.Println(arr, yArr)

	// The slice we create for our function below:
	var slice1 = []float32{123.66, 77.45, 23.49}
	fmt.Println("Before we use our function: ", slice1)
	changeFirst(slice1)
	fmt.Println("After we use our function: ", slice1)

}

// a Function that excepts mutable datatype:
func changeFirst(slice []float32) {
	slice[0] = 121.45
}
