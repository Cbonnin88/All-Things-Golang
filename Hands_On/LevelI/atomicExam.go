package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() { //nolint:typecheck
	fmt.Println("CPUs in the Beginning:",runtime.NumCPU())
	fmt.Println("goRoutines in the beginning",runtime.NumGoroutine())

	var counter int64
	gr := 120
	var wg sync.WaitGroup
	wg.Add(gr)

	for i := 0; i < gr; i ++ {
		go func() {

			at := atomic.AddInt64(&counter,2)
			fmt.Println(at)
			wg.Done()
		}()

	}
	wg.Wait()
	fmt.Println("end value:",counter)



}
