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
	const gr = 200
	var wg sync.WaitGroup
	wg.Add(gr)

	var mu sync.Mutex

	for i := 0; i < gr; i ++ {
		go func() {
			mu.Lock()
			v := counter
			v++
			counter = v
			fmt.Println(counter)
			mu.Unlock()
			wg.Done()
		}()

	}
	wg.Wait()
	fmt.Println("end value:",counter)



}
