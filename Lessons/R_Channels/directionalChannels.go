package main

import "fmt"

func main() { //nolint:typecheck
	/*
	Channels can be Bi-directional, send-only or receive only
	send-only : chan <- int
	receive-only: <- chan int
	*/

	so := make(chan string, 1)
	sendOnly(so)
	fmt.Println(<-so)

	ro := make(chan string, 1)
	ro <- "I am receive-only brah !!!"
	receiveOnly(ro)

	fmt.Println("")
	fmt.Printf("Send-only is of type: %T\n",so)
	fmt.Printf("Receive only is of type: %T\n",ro)


}

// Send-only Channel
func sendOnly(so chan <- string) {
	so <- "I am send-only dude!!!"
}

// Receive-only Channel
func receiveOnly(ro <- chan string){
	fmt.Println(<-ro)
}
