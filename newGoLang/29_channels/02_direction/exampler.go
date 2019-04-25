package main

import "fmt"

func main() {

	c := make(chan int, 2)

	go foo(c)
	go bar(c)

	fmt.Println("About to exit")

}

func foo(c chan<- int) {
	c <- 50
} //send only channel

func bar(c <-chan int) {
	fmt.Println(<-c)
} //receive only channel
