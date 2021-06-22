package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Our OS:", runtime.GOOS)
	fmt.Println("Our Architecture:", runtime.GOARCH)
	fmt.Println("Our CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:",runtime.NumGoroutine())

}
