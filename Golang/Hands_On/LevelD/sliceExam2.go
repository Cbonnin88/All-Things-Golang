package main

import "fmt"

func main() { //nolint:typecheck
	slice := []int{72,73,74,75,76,77,78,79,80,81}
	slice = append(slice,82)
	fmt.Println(slice)
	fmt.Println("")

	slice = append(slice,83,84,85)
	fmt.Println(slice)
	fmt.Println("")

	slice2 :=[]int{86,87,88,89,90}
	slice = append(slice,slice2...)
	fmt.Println(slice)


}
