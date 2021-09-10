package main

import "fmt"

func main() {
	x := []int{22, 23, 24, 25, 26, 27, 28, 29, 30, 31}

	y := append(x[0:3], x[6:10]...)
	fmt.Println(y)

}
