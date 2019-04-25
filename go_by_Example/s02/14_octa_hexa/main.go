// File name: ...\s02\14_octa_hexa\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import "fmt"

func main() {
	//decimal literal declaration
	d := 125
	fmt.Printf("%8d %#8o %#8x \t %08b\n", d, d, d, d)

	//octa literal - 0 added on the left side
	oc := 0175
	fmt.Printf("%8d %#08o %#8x \t %08b\n", oc, oc, oc, oc)

	//hex literal - 0x added on the left side
	x := 0x7d
	fmt.Printf("%8d %#8o %#8x \t %08b\n", x, x, x, x)

	// b := 01111101 // binary literals are not supported in Go
}
