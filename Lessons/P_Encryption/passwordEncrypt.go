package main

import (
	"fmt"


)


func main() {
	pwd := `catfood1989`
	bs, err := GenerateFromPassword([]byte(pwd), MinCost)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(pwd)
	fmt.Println(bs)
	fmt.Println("")

	loginPwd := `catfood1989`
	err = CompareHashAndPassword(bs, []byte(loginPwd))
	if err != nil {
		fmt.Println("INCORRECT PASSWORD")
		return
	}
	fmt.Println("You have successfully logged in")
}
