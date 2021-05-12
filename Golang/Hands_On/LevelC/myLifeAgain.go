package main

import "fmt"

func main() {
	fmt.Println("The years I've been alive, counting backwards:")
	y := 2021
	for {
		if y < 1988 {
			break
		}
		fmt.Println(y)
		y--
	}

}
