package main

import "fmt"

func main() {
	d := 10
	e := 20
	f := 21009

	fmt.Println(tripleNum(d))
	fmt.Println(tripleNum(e))
	fmt.Println(tripleNum(f))
	jobOffer()

}

func tripleNum(num int) int {
	return num * 3
}

func jobOffer() {
	fmt.Println("Congrats, you're hired")
}
