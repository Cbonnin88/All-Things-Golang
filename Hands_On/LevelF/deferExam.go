package main

import "fmt"

func main() {
	poem()
	poem2()
	poem3()
}

func poem() {
	defer func() {
		fmt.Println("Who swore up and down that he could dance,")
	}()
	fmt.Println("There once was a man from France,")
}

func poem2() {
	defer fmt.Println("He fell under a trance,")
	fmt.Println("And with just one glance,")
}

func poem3() {
	defer fmt.Println("And at last he had found romance,")
	fmt.Println("for a beautiful lady had come by no mischance,")
}
