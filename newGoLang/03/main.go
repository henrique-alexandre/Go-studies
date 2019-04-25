package main

import "fmt"

func main() {

	x := 42
	fmt.Println(x)
	x = 99
	fmt.Println(x)
	y := "James Bond"
	fmt.Println(y)

	fmt.Printf("%T\n %T\n", y, x)

	s := fmt.Sprintf("%#x\t %b\t %x\t %v", x, x, x, x)
	fmt.Println(s)

	//statements are made up of expressions.
	//primitive and composite data types

}
