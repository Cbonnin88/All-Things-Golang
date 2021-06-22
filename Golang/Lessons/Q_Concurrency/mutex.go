package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() { //nolint:typecheck
	fmt.Println("CPUs being used:",runtime.NumCPU())
	fmt.Println("GoRoutines that currently exist:", runtime.NumGoroutine())

	counter := 0
	const goRoutines = 220
	var wg sync.WaitGroup
	wg.Add(goRoutines)

	var mu sync.Mutex // here we create our Mutex variable

	for i := 0; i < goRoutines; i++ {
		go func(){
			mu.Lock() // here we lock our 'counter' variable
			v := counter
			runtime.Gosched() // Gosched yields the CPU allowing other goroutines to run
			v++
			counter = v
			mu.Unlock()
			wg.Done()
		}()
		fmt.Println("GoRoutines:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("GoRoutines:", runtime.NumGoroutine())
	fmt.Println("count:",counter)
	fmt.Println("here we have no race conditions")
}
