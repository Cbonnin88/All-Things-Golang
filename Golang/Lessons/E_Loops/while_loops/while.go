package main

import "fmt"

func main() {
	num := 1
	for num < 5 {
		num += 3
	}
	fmt.Println(num)

	// Using Continue:
	for x := 0; x <= 1000; x++ {
		if x != 0 && x%3 == 0 && x%7 == 0 && x%9 == 0 {
			fmt.Println(x)
			continue
		}
	}
	// Using Break:
	for num := 0; num <= 1000; num++ {
		if num != 0 && num%2 == 0 && num%4 == 0 && num%6 == 0 {
			fmt.Println(num)
			break
		}
	}

	fmt.Println("While Loop:")
	i := 0
	for i < 5 {
		i += 2
		fmt.Println(i)
	}
}
