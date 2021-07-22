package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() { //nolint:typecheck
	fmt.Println("CPUs being used:", runtime.NumCPU())
	fmt.Println("GoRoutines that currently exist:", runtime.NumGoroutine())

	var counter int64
	const goRoutines = 220
	var wg sync.WaitGroup
	wg.Add(goRoutines)



	for i := 0; i < goRoutines; i++ {
		go func() {
			atomic.AddInt64(&counter, 1) // Here we use the atomic package but giving the address to our counter variable and by how much to increase
			fmt.Println("Counter\t",atomic.LoadInt64(&counter))
			runtime.Gosched()
			wg.Done()
		}()
		fmt.Println("GoRoutines:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("GoRoutines:", runtime.NumGoroutine())
	fmt.Println("count:", counter)

}
