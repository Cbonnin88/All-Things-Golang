package main

import "fmt"

// Package level Scope
var x string = "Rest While you can because I will find you, and I will break you!!!"

func main() {
	fmt.Println(x)
	alcina()
}

func alcina() {
	fmt.Println(x)
}
