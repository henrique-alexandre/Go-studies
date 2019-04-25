package main

import "fmt"

func main() {

	fmt.Println(foo(10))

	n := 0

	for i := 10; i > 0; i-- {

		n = n + i

	}

	fmt.Println(n)

	//for i := 1; i < 50; i++ {
	//	if i %3 == 0 {
	//		defer fmt.Print(i, " ")
	//	}
	//
	//}
}

var f int = 0

func foo(x int) int {

	if x == 0 {
		return f
	}

	f += x
	x--
	return foo(x)

	//return x + foo(x-1)

}
