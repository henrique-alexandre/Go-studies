package main

import "fmt"

const o = 42
const oo = "Big Bang"

const (
	primeira int     = 1000
	segunda  float64 = 200.20
	terceira string  = "Ahava"
)

func main() {

	a := "Henrique"
	b := []byte(a)
	fmt.Println(b)
	fmt.Printf("%T\n", b)

	for i := 0; i < len(b); i++ {

		fmt.Printf("%#x ", b[i])
	}

	for i, v := range b {
		fmt.Println(i, v)
	}
	for i, v := range b {
		fmt.Printf("at index position %d we have hex %#x\n", i, v)
	}

}
