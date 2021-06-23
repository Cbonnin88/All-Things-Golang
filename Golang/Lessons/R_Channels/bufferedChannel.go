package main

import "fmt"

func main() { //nolint:typecheck

	/*
	A Buffered Channel is a channel that accepts a limited number of values without a corresponding receiver for those values
	*/
	ch := make(chan int, 2) // Here we create a buffer channel that will allow two values to sit in

	ch <- 50
	ch <- 75

	fmt.Println("First value:",<-ch)
	fmt.Println("Second Value:",<-ch)

	// Example 2

	ch2 := make(chan string, 3)

	ch2 <- "Here"
	ch2 <- "Is"
	ch2 <- "Something"
	fmt.Println(<-ch2,<-ch2,<-ch2)
}
