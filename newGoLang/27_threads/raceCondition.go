package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("CPUs: ", runtime.NumCPU())
	//fmt.Println("Number of routines: ", runtime.NumGoroutine())

	counter := 0

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)
	//Mutex is used to prevent race conditions
	var mu sync.Mutex

	for i := 0; i < gs; i++ {

		go func() {
			mu.Lock()
			v := counter
			v++
			counter = v
			mu.Unlock()
			wg.Done()
		}()
		fmt.Println("Number of routines: ", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("Hello", counter)
	fmt.Println("Number of routines: ", runtime.NumGoroutine())

}
