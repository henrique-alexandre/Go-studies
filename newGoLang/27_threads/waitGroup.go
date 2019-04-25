package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {

	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.NumGoroutine())

	wg.Add(2)
	go foo()

	go bar()

	wg.Wait()
}

func foo() {
	for i := 0; i <= 100; i++ {
		fmt.Println("foo: ", i)
	}
	wg.Done()
}
func bar() {
	for i := 0; i <= 100; i++ {
		fmt.Println("bar: ", i)
	}
	wg.Done()
}
