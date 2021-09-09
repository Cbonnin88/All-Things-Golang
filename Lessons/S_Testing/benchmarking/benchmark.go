package main

import (
	"fmt"
)

func main() {
	fmt.Println(expression("Ethan Winters"))
}

// Expressions takes in a string value and returns a string value
func expression(s string) string {
	return fmt.Sprint("Rest while you can, ", s, " "+
		",because I will hunt you, and I will break you !!!")
}
