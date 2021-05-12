package main

import "fmt"

func main() { //nolint:typecheck
	fmt.Println("A switch statement with no statement")

	switch {
	case false:
		fmt.Println("Only a Sith deals in lies")
	case true:
		fmt.Println("The truth will sent you free")
	}
}
