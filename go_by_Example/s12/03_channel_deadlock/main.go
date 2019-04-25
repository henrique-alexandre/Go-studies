// File name: ...\s12\03_channel_deadlock\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import "fmt"

func main() {
	c := make(chan string)
	//c <- "No one likes my channel!" //fatal error: all goroutines are asleep - deadlock!
	go sex (c)
	fmt.Println(<-c)
}

//A Channel is a vehicle that facilitates communication between GoRoutines.

func sex (c chan string) {
	c<-"oi"
	close(c)
}