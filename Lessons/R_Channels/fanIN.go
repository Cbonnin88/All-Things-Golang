package main

import (
	"fmt"
	"sync"
)

func main() {

	/*
		Fan in : Take values from many channels and putting those values onto one
		channel
	*/

	even := make(chan int)
	odd := make(chan int)
	fanin := make(chan int)

	go sendFanIn(even, odd)
	go receiveFanIn(even, odd, fanin)

	// here we range our fan in
	for v := range fanin {
		fmt.Println(v)
	}

	fmt.Println("about to exit")
}

func receiveFanIn(ev, od <-chan int, fanin chan<- int) {
	var wg sync.WaitGroup
	wg.Add(2) // here we launch our wait groups, one for even, one for odd

	go func() {
		for v := range ev {
			fanin <- v
		}
		wg.Done()
	}()
	go func() {
		for v := range od {
			fanin <- v
		}
		wg.Done()
	}()
	wg.Wait()
	close(fanin)

}

func sendFanIn(ev, od chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			ev <- i
		} else {
			od <- i
		}
	}
	close(ev)
	close(od)
}
