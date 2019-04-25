package main

import "fmt"

type cat int

func main() {

	a := 200
	var q cat = 10

	fmt.Printf("%d\t %b\t %x\n", a, a, a)

	a = int(q)

	x := q << 2
	fmt.Printf("%b\t %d\t %x\n", x, x, x)

	s := `String "something else" nice`
	fmt.Println(s)

}
