package main

import "fmt"

func main() {

	e := fuu()
	fmt.Println(e())
	//fmt.Println(foo())
	//fmt.Println(bar())

	//anonymous func
	f := func() {
		for i := 0; i <= 100; i++ {
			fmt.Println(i)
		}
	}
	f()

	fmt.Printf("%T", f)

	a := foo()
	b, d := bar()

	fmt.Println(a, b)
	fmt.Println(d)

}

func foo() int {
	return 10
}

func bar() (n int, s string) {
	return 50, "Henrique"

}

func fuu() func() int {
	return func() int {
		return 42

	}

}
