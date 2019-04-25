package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	f, err := os.Open("Teste")
	if err != nil {
		fmt.Println(err)
	}

	sb := make([]byte, 100)

	f.Read(sb)
	s := string(sb)
	fmt.Println(s)

	x, e := ioutil.ReadFile("Teste")
	if err != nil {
		fmt.Println(e)
	}

	t := string(x)
	fmt.Println(t)

	copia, _ := os.Open(os.Args[1])
	fmt.Println(copia)
	fmt.Println(*copia)

	io.Copy(os.Stdout, copia)

	fmt.Println()

}
