package main

import "fmt"

func main() {
	// using an 'if' statement and a the 'break' keyword
	x := 1
	for {
		if x > 9 {
			// We are 'breaking' our loop here, where x is greater than 9:
			break
		}
		fmt.Println(x)
		x++
	}
	fmt.Println("All done.")
	fmt.Println("")
	fmt.Println("counting by fours:")

	// a for loop with a 'continue' keyword
	num := 1
	for {
		num++ // we place our increment at the top
		if num > 100 {
			break // will break when it hits a number greater than 100
		}
		if num%4 != 0 {
			continue
		}
		fmt.Println(num)
	}

}
