package main

import (
	"fmt"
)

func main() {
	fmt.Println(minus(2, 10))

}

func minus(num ...int) int {
	var subtr int
	for _, v := range num {
		subtr = v - subtr
	}
	return subtr
}
