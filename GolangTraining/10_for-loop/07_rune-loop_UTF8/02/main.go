package main

import "fmt"

func main() {
	for i := 100; i <= 300; i++ {
		fmt.Printf("%v - %v - %v \n", i, string(i), []byte(string(i)))
	}
}
