package main

import (
	"fmt"
	"runtime"
	"sync"
)

var c int
var w sync.WaitGroup
var m sync.Mutex

func main() {

	w.Add(2)

	go func() {
		runtime.Gosched()
		m.Lock()
		c++
		m.Unlock()
		w.Done()
	}()
	go func() {
		runtime.Gosched()
		m.Lock()
		c++
		m.Unlock()
		w.Done()
	}()

	fmt.Println(runtime.NumGoroutine())
	w.Wait()

	fmt.Println(c)

	fmt.Println(runtime.NumGoroutine())

}
