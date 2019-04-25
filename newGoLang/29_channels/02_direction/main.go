package main

import "fmt"

//A channel might be bidirectional or unidirectional.
//Bidirectional: sends and receives
//Unidirectional: either receives or sends

func main() {

	c := make(chan int, 2)
	cr := make(<-chan int, 2) // receive
	cs := make(chan<- int, 2) // send

	c <- 40
	c <- 50

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Printf("%T\n", c)
	fmt.Printf("%T\n", cr)
	fmt.Printf("%T\n", cs)

}
