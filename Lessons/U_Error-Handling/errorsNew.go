package main

import (
	"errors"
	"fmt"
	"log"
)

var errorMessage = errors.New("square root of a negative number") // Here we create or error variable

func main() {
	// The errors.New() method can be used to create new errors and takes the error message as its only parameter.
	fmt.Printf("errorMessage is of type %T\n", errorMessage)
	_, err := sqrt(-20)
	if err != nil {
		log.Fatalln(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errorMessage
	}
	return 32, nil
}
