package main

import "fmt"

func main() { //nolint:typecheck
	ch := make(chan int)
	go func() {
		ch <- doSomething(10) // putting the value of our function into a channel
	}()
	fmt.Println(<-ch) // Here we pull the value off the channel and print it
}

func doSomething (x int) int {
	// Do something , already built
	return x * 5
}
