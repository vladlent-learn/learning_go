package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())

	counter := 0
	const gs = 100

	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			v := counter

			runtime.Gosched()

			v++
			counter = v

			fmt.Println("Goroutines:", runtime.NumGoroutine())
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("counter:", counter)

}
