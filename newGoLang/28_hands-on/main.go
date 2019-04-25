package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {

	wg.Add(2)

	go func() {
		fmt.Println("Henrique")
		wg.Done()
	}()

	go func() {
		fmt.Println("Alexandre")
		wg.Done()
	}()

	fmt.Println(runtime.NumGoroutine())

	wg.Wait()

	fmt.Println("finished")

	fmt.Println(runtime.GOARCH)
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.GOROOT())
	s := runtime.GOOS
	fmt.Println(s)

}
