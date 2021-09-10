package main

import "fmt"

func main() {
	fmt.Println("A switch statement with no statement")

	switch {
	case false:
		fmt.Println("You stupid man-thing!!!")
	case true:
		fmt.Println("Rest while you can, because I will hunt you and I will break you!!!")
	}
}
