package main

import "fmt"

func main() { //nolint:typecheck

	// nesting a Loop
	fmt.Println("Nesting loops")
	for i := 0; i <= 20; i++ {
		fmt.Printf("The outer loop: %d\n", i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t\tThe inner Loop: %d\n", j)
		}

	}
}
