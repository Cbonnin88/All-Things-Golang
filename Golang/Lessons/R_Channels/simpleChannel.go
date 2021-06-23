package main

import (
	"fmt"
	"runtime"
)

func main() { //nolint:typecheck

	c := make(chan int) // here we create a channel in which we can place integers, like maps, channels are allocated with 'make'
	go func(){
		c <-50 // Here we put '50' onto our channel
	}()
	fmt.Println(<-c) // here we tell it to take '50' off our channel
	fmt.Println("Goroutines:",runtime.NumGoroutine())
	fmt.Printf("Our channel is of type: %T\n",c)
}
