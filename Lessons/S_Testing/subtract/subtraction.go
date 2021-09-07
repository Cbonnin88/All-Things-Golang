package main

import (
	"fmt"
)

func main() {
	fmt.Println(minus(2, 10))

}

func minus(num ...int) int {
	subtr := 0
	for _, v := range num {
		subtr = v - subtr
	}
	return subtr
}
