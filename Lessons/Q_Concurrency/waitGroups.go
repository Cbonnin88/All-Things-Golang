package main

import (
	"fmt"
	"runtime"
	"sort"
	"sync"
)

var wg sync.WaitGroup

func main() { //nolint:typecheck

	fmt.Println("OS:", runtime.GOOS)                   // Gives us the Operating System
	fmt.Println("Arch:", runtime.GOARCH)               // Gives us the Architecture structure
	fmt.Println("CPUs:", runtime.NumCPU())             // Gives us the number of CPU's
	fmt.Println("Goroutines:", runtime.NumGoroutine()) // Gives us the number of Go-routines

	wg.Add(1)  // here we add the number of programs we want it to wait for
	go sort1() // here we launch a new Go routine
	sort2()

	fmt.Println("CPUs in use:", runtime.NumCPU())
	fmt.Println("Goroutines currently active:", runtime.NumGoroutine())
	wg.Wait()
}

func sort1() {
	slice := []int{2, 5, 1, 3, 7, 9, 8, 6, 10, 4}
	sort.Ints(slice)
	fmt.Println(slice)
	wg.Done() // when this program is finished running, this will be removed
}

func sort2() {
	names := []string{"Barbara", "Zaddy", "Alcina", "Fred", "Melia"}
	sort.Strings(names)
	fmt.Println(names)
}
