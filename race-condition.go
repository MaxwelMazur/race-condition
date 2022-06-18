package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("CPUs: ", runtime.NumCPU())
	fmt.Println("Goroutines: ", runtime.NumGoroutine())

	count := 0
	totalGoroutines := 1000

	var wg sync.WaitGroup
	wg.Add(totalGoroutines)

	for i := 0; i < totalGoroutines; i++ {
		go func() {
			v := count
			runtime.Gosched()
			v++
			count = v
			wg.Done()
		}()
		fmt.Println("Goroutines: ", runtime.NumGoroutine())
	}

	wg.Wait()

	fmt.Println("========================")
	fmt.Println("Goroutines: ", runtime.NumGoroutine())
	fmt.Println("Value final: ", count)
}
