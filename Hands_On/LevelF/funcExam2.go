package main

import "fmt"

func main() { //nolint:typecheck

	dimitrescu := func() {
		fmt.Println("You stupid Man-Thing!!!")
	}
	dimitrescu()

	dimitrescu2 := func(word string) {
		fmt.Println(word)

	}
	dimitrescu2("You won't live long, even if you run!!!")

}
