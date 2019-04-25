// File name: ...\s02\13_bitwise_operators\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import "fmt"

func main() {

	var x uint8 = 2


	fmt.Printf("%8d %#8o %#8x \t %08b\n", x, x, x, x)

	x = x << 1
	fmt.Printf("%8d %#8o %#8x \t %08b\n", x, x, x, x)

	x = x >> 2
	fmt.Printf("%8d %#8o %#8x \t %08b\n", x, x, x, x)

	// 00000100 --> 4
	// 00000010 --> 2
	// 00000110 --> 6 = 4+2

	x = 4 | 2
	fmt.Printf("%8d %#8o %#8x \t %08b\n", x, x, x, x)

	// 00000100 --> 4
	// 00000010 --> 2
	// 00000000 --> 0
	x = 4 & 2
	fmt.Printf("%8d %#8o %#8x \t %08b\n", x, x, x, x)

	x = 1 ^ 4	//compiler error: constant -5 overflows uint8
	fmt.Printf("%8d %#8o %#8x \t %08b\n", x, x, x, x)


	// 00000101 --> 5
	// 11111111 --> 255
	// 11111010 --> 250
	fmt.Printf("%8d %#8o %#8x \t %08b\n", ^x, ^x, ^x, ^x)

	// 00000100 --> 4
	// 00000001 --> 1
	// 00000101 --> 5
	y := 1 ^ 4
	fmt.Printf("%8d %#8o %#8x \t %08b\t %T\n", y, y, y, y, y)



}
