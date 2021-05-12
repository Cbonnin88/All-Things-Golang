package main

import "fmt"

func main() { //nolint:typecheck
	fmt.Println("Printing the alphabet")
	for i := 65; i <= 90; i++ {
		fmt.Println(i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t%#U\n", i)
		}
	}

}
