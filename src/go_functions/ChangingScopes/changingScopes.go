package main

import "fmt"

func multiply10(numMult *int) {
	*numMult = 10
}

func main() {
	multiply := 20
	multiply10(&multiply)
	fmt.Println(multiply * 10)

}
