package main

import "fmt"

func main() {
	ref := reference
	// ref() is the same as reference(), it calls our function
	ref()

	// calling a function with a parameter
	ref2 := reference2
	ref2("I have been referenced again")

	// creating and calling a function inside our main()
	words := func() {
		fmt.Println("I am a function within a function")
	}
	words()

	// creating a function with a return
	math := func(x int) float32 {
		return float32(x) * 5.9
	}(9)
	fmt.Println(math)

	// Calling our nested function
	math3 := func(y float32) float32 {
		return y / 21.54
	}
	math2(math3)

	// calling another nested function
	words2 := func(letter string) string {
		return letter + "is how I like it"
	}
	words3(words2)

	// returning our function in the main
	returnFunc("I am returning a function")()
	goodbye := returnFunc("I am learning Golang")
	goodbye()
}

// 1.simple referencing
func reference() {
	fmt.Println("I have been referenced")
}

// 2. referencing with a parameter
func reference2(w string) {
	fmt.Println(w)
}

// 3. Passing a function into another function
func math2(myFunc func(float32) float32) {
	fmt.Println(myFunc(12.5))
}

// Passing another function into a string function
func words3(myFunc func(string) string) {
	fmt.Println(myFunc("Baked "))
}

// returning a function
func returnFunc(x string) func() { // also known as a function closure
	return func() {
		fmt.Println(x)
	}
}
