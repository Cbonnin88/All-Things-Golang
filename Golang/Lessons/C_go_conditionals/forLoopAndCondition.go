package main

import "fmt"

func main() {
	for i := 1; i < 250; i++ { // our Loop
		if i%4 == 0 { // our condition and modulo
			fmt.Println(i)
		}
	}

}
