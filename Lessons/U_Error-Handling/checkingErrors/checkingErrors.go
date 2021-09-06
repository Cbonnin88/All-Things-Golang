package main

import "fmt"

func main() {
	nb, err := fmt.Println("Hi there")
	if err != nil { // This means that there is an error, == means that there's no error
		fmt.Println(err)
	}
	fmt.Println(nb)
}
