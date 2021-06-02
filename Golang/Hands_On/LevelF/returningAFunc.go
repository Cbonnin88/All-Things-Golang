package main

import "fmt"

func main() { //nolint:typecheck

	letsReturn := returnMe()
	fmt.Println(letsReturn())
	fmt.Printf("letsReturn() is of type:%T",letsReturn())

}

func returnMe() func() float32 {
		return func() float32 {
			return 32.6
		}
}
