package main

import "fmt"

func main() { //nolint:typecheck
	expression := func(){
		fmt.Println("I am a func expression with no parameters or arguments")
	}
	expression()

	expression2 := func(y string) {
		fmt.Println("I am a func expression with one parameter and an argument of:",y)
	}
	expression2("func expressions gurl")

}
