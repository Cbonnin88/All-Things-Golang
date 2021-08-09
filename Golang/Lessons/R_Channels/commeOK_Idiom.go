package main

import "fmt"

func main () { //nolint:typecheck

	// Using 'Comma ok Idiom'

	c := make(chan string)

	go func() {
		c <- "Well, Well, Well Ethan Winters"
		close(c)
	}()

	v, ok := <-c
	fmt.Println(v, ok)

	v, ok = <-c
	fmt.Println(v,ok)
}