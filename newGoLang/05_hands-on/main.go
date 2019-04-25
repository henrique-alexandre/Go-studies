package main

import "fmt"

type cat int

var h int
var e string
var n bool

func main() {

	x := 42
	y := "James Bond"
	z := true

	fmt.Println(x, y, z)
	fmt.Println(h, e, n)
	e = "Henrique"

	s := fmt.Sprintf("%v\t %v\t %v\t", h, e, n)
	fmt.Println(s)

	var q cat
	fmt.Printf("%v \t %T \n", q, q)
	q = 42
	fmt.Println(q)

	u := int(q)
	fmt.Print(u)

}
