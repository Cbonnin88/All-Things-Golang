package main

import "fmt"

func main() { //nolint:typecheck
	c := make(chan float64)

	// Send-only
	go sending(c)

	// Receive-only
	receiving(c)

	fmt.Println("about to exit")


}
func sending(c chan <- float64){
	c <- 123.55
}

func receiving(c <- chan float64){
	fmt.Println(<-c)
}
