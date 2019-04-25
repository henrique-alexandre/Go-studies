package main

import (
	"fmt"
	"runtime"
)

var x bool

func main() {

	a := "Henrique"
	b := "Henrique"

	fmt.Println(a == b)

	fmt.Println(runtime.GOARCH)
	fmt.Println(runtime.GOOS)

}
