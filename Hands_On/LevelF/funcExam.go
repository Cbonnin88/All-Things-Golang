package main

import "fmt"

func main() { //nolint:typecheck
	slice :=[]int {2,4,6,8,10,12}
	multiply(slice...)
	fmt.Println("")

	slice2 :=[]float32 {2.2,3.3,4.4,5.5,6.6}
	divide(slice2)

}

func multiply(sum... int) int {
	fmt.Println(sum)

	multiply := 1
	for i, v := range sum {
		multiply *= v
		fmt.Println("Index:",i,"we are multiplying",v,"to the total, which is now",multiply)
	}
	fmt.Println("The total is",multiply)
	return multiply
}

func divide(sum2 []float32) float32 {
	divide := 1.1
	for index, value := range sum2 {
		divide /= .2
		fmt.Println("Index:",index,"we are dividing",value,"to the total, which is now",divide)
	}
	fmt.Println("The total is",divide)
	return float32(divide)
}
