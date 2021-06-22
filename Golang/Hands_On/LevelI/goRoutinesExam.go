package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() { //nolint:typecheck
	fmt.Println("CPUs in the Beginning:",runtime.NumCPU())
	fmt.Println("goRoutines in the beginning",runtime.NumGoroutine())

	counter := 0
	const gr = 100
	var wg sync.WaitGroup
	wg.Add(gr)

	for i := 0; i < gr; i ++ {
		go func() {
			v := counter
			runtime.Gosched()
			v++
			counter = v
			fmt.Println(counter)
			wg.Done()
		}()

	}
	wg.Wait()
	fmt.Println("end value:",counter)
	fmt.Println("Found one data race")


}
