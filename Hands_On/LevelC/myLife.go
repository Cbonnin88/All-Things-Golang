package main

import "fmt"

func main() {
	fmt.Println("All the years I've been alive:")
	bd := 1988
	for bd <= 2021 {
		fmt.Println(bd)
		bd++
	}

}
