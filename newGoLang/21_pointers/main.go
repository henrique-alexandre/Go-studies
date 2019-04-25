package main

//Pass by value!

import "fmt"

func main() {

	a := 32
	fmt.Println(a)
	fmt.Println(&a) //&ampersand gives the address
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a)

	var b *int = &a //asterisk works in this line as a pointer to a value
	fmt.Println(b)
	fmt.Println(*b) //asterisk works in this line as a de-referencing element, giving the value stored in the address
	fmt.Println(*&a)
}
