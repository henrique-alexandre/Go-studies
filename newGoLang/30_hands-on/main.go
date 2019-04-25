package main

import "fmt"

func main() {

	c := make(chan int)
	//c := make(chan int, 1) //solve the problem without creating an anonymous func

	go func() { c <- 42 }()

	fmt.Println(<-c)

	cs := make(chan int)

	go func() {
		cs <- 42
	}()
	fmt.Println(<-cs)

	fmt.Printf("------\n")
	fmt.Printf("cs\t%T\n", cs)

}
