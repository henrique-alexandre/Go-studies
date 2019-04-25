package main

import "fmt"

func main() {

	var myGreeting map[string]string
	fmt.Println(myGreeting)
	fmt.Println(myGreeting == nil)

	n := map[int]string{0: "Henrique", 1: "Cilene"}
	fmt.Println(n)

	//mySlice := make([]int, 3, 6)
	//fmt.Println(mySlice[2])
}

// add these lines:
/*
	myGreeting["Tim"] = "Good morning."
	myGreeting["Jenny"] = "Bonjour."
*/
// and you will get this:
// panic: assignment to entry in nil map

//Maps are unordered group of elements of one type, called the element type.
//It's indexed by a set of unique keys of another type, called the key type.
//Maps don't have append functions.
