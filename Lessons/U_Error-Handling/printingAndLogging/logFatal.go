package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	defer foo()
	_, err := os.Open("dogs")
	if err != nil {
		log.Fatalln("CANNOT FIND FILE") //
	}
}

func foo() {
	fmt.Println("When os.Exit() is called, deferred functions will not run")
}
