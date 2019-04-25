package main

import "fmt"

func main() {

	c := make(chan int)

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)

	d := make(chan int, 1)

	d <- 40

	fmt.Println(<-d)

}
