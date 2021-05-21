package main

import "fmt"

func main() {
	var money [5]int
	money[0] = 100
	money[1] = 250
	money[2] = 350
	money[3] = 450
	money[4] = 550

	for i, element := range money {
		fmt.Println(i,element)
	}
	fmt.Printf("Our array type: %T",money)
}
