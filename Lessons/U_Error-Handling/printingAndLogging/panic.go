package main

import "os"

func main() {
	_, err := os.Open("doctor")
	if err != nil {
		panic(err) // panic stops normal execution of goroutines
	}

}
