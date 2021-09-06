package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.Open("pie.txt")
	if err != nil {
		fmt.Println("FILE DOES NOT EXIST") // prints to standard out
	}
}
