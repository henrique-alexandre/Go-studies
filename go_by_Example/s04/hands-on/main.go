package main

import "fmt"

// ASSIGNMENT: implementing basic operations with slice.
// Use the slicing techniques you've learned so far to implement
// the following basic operations with slices:
// Append (a=[1 2 3] & b=[4 5 6 7] -> c=[1 2 3 4 5 6 7])
// Copy (a=[1 2 3] -> b=[1 2 3])
// Cut (a=[1 2 3 4 5 6 7] -> a=[1 2 5 6 7])
// Delete (with preserving order)    (a=[1 2 5 6 7] -> a=[1 5 6 7])
// Delete (without preserving order) (a=[1 2 3 4 5 6 7] -> a=[1 2 7 4 5 6])
// Expand (a=[1 2 7 4 5 6] -> a=[1 2 0 0 0 0 7 4 5 6])
// Extend (a=[1 2 0 0 0 0 7 4 5 6] -> a=[1 2 0 0 0 0 7 4 5 6 0 0 0])
// Insert (a=[1 2 0 0 0 0 7 4 5 6 0 0 0] -> a=[1 2 0 9 10 0 0 0 7 4 5 6 0 0 0])
//
// Ref: https://github.com/golang/go/wiki/SliceTricks

func main() {

	a := []int{1,2,3,4,5}
	b := []int{6,7,8,9,10}

	a = append(a, b...)
	fmt.Println(a)

	c := make([]int, 10, 20)
	copy(c, a)
	c = append(c, 50,60,70)
	fmt.Println(c)

	d := append([]int(nil), a...)
	fmt.Println(d)

	d = append(c[:2], c[6:]...)
	fmt.Println(d)





}
