package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("Beginning CPUs:", runtime.NumCPU())
	fmt.Println("Beginning goRoutines:",runtime.NumGoroutine())

	var wg sync.WaitGroup
	wg.Add(2)

	go func(word string) string{
		return word + "cat"
	}("house")
	wg.Done()


	go func(x int) int {
		return x + 5
	}(10)
	wg.Done()

	fmt.Println("mid CPU:",runtime.NumCPU())
	fmt.Println("mid goRoutine:",runtime.NumGoroutine())

	wg.Wait()

	fmt.Println("about to exit")
	fmt.Println("end CPUs:",runtime.NumCPU())
	fmt.Println("end goRoutines:",runtime.NumGoroutine())

}









