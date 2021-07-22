package main

import "fmt"

func main() { //nolint:typecheck
	/*
	SELECT lets you wait on multiple channel operations
	*/
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	// send
	go send(even, odd, quit)

	// receive
	receive(even, odd, quit)

	fmt.Println("about to exit")
}

func receive(ev, od, qu <-chan int){
	for  {
		select { // Here we use 'select' to await all three values simultaneously, printing each one as it arrives
			case v := <- ev:
				fmt.Println("From the even channel:",v) // will receive our even values
		case v := <- od:
			fmt.Println("From the odd channel",v) // will receive our odd values
		case i, ok := <- qu: // Here have our 'ok' idiom using an int
			if !ok {
				fmt.Println("from the comme ok:",i, ok)
				return
			} else {
				fmt.Println("from the comma ok:",i)
			}
		}
	}
}

func send(ev, od, qu chan<- int){
	for i := 0; i < 10; i++ { // here we create our loop to get our even and odd values using a modulus '%'
		if i % 2 == 0 {
			ev <- i // our even values
		} else {
			od <- i // our odd values
		}
	}
	close(qu)
}
