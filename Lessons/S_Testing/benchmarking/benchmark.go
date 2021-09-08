package main

import (
	"fmt"
)

func main() {
	fmt.Println(expression("Ethan Winters"))
}

func expression(s string) string {
	return fmt.Sprint("Rest while you can, ", s, " "+
		",because I will hunt you, and I will break you !!!")
}
