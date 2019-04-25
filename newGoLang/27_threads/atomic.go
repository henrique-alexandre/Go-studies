package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

//package Atomic requires an int64

func main() {

	const gs = 100
	var counter int64
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
		}()
		fmt.Println(atomic.LoadInt64(&counter))
		wg.Done()

	}
	wg.Wait()
	fmt.Println(counter)
}
