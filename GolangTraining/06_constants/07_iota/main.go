package main

import "fmt"

const (
	//B  = 1 << (iota * 10) // 1
	KB = 1 << (iota * 10) // 1 << (1 * 10)
	MB = 1 << (iota * 10) // 1 << (2 * 10)
	GB = 1 << (iota * 10) // 1 << (3 * 10)
	TB = 1 << (iota * 10) // 1 << (4 * 10)
)

//const (
//	uB  = 1 << (iota * 10) // 1
//	uKB = 1 << (iota * 10) // 1 << (1 * 10)
//	uMB = 1 << (iota * 10) // 1 << (2 * 10)
//	uGB = 1 << (iota * 10) // 1 << (3 * 10)
//	uTB = 1 << (iota * 10) // 1 << (4 * 10)
//)
//
//func blah() {
//	_ = 00000000000000000000000000000001
//	_ = 00000000000000000000010000000000
//
//	_ = 00000000000000000000010000000000
//	_ = 00000000000100000000000000000000
//
//	_ = 01010100001000111010101
//	_ = 10000100011101010100000 // is the above number << 5
//	_ = 10000100011101010100000 // is the above number << 5
//
//}

func main() {
	fmt.Println("binary\t\tdecimal")
	//fmt.Printf("%b\t", B)
	//fmt.Printf("%b\t", B)
	fmt.Printf("%b\t", KB)
	fmt.Printf("%d\n", KB)
	fmt.Printf("%b\t", MB)
	fmt.Printf("%d\n", MB)
	fmt.Printf("%b\t", GB)
	fmt.Printf("%d\n", GB)
	fmt.Printf("%b\t", TB)
	fmt.Printf("%d\n", TB)
}
