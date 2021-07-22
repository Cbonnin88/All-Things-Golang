package main

import "fmt"

func main() { //nolint:typecheck
	c := make(chan float64)

	// Send-only
	go func() {
		for i := 0.0; i < 100.00; i++ {
			c <- i
		}
		close(c)
	}()

	// Receive-only
	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("about to exit")


}




