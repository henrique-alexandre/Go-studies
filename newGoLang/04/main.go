package main

import "fmt"

type hotdog int

func main() {

	var a = 42
	var b hotdog = 43

	fmt.Print(a, "is of type ")
	fmt.Printf("%T\n", a)
	fmt.Print(b, "is of type ")
	fmt.Printf("%T\n", b)

	//this is CONVERSION, not Casting!
	a = int(b)
	fmt.Print(a)

}
