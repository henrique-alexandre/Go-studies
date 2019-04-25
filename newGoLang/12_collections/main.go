package main

import "fmt"

//an Array is a building block, a numbered sequence of elements of a single type, called element type.
// The number of elements is called length and is never negative. The length is part of the Array type.

func main() {

	//example of array
	var x [5]int
	x[3] = 42
	fmt.Println(x)
	fmt.Println(len(x))

	//Slices are preferred over Arrays
	//composite literal: x := type{values}. It's an expression that creates a new instance each time it is evaluated.

	q := []int{10, 20, 40, 80}
	fmt.Println(q, "		", len(q))

	q = append(q, 100, 200, 300)

	fmt.Println(q[1:])
	fmt.Println(q[:3])
	fmt.Println(q[1:3]) //it doesn't bring the last element

	for i, v := range q {
		fmt.Println("position ", i, "is ", v)
	}

	for i := 0; i < len(q); i++ {
		fmt.Println(q[i])
	}

	h := []int{2, 4, 8, 16, 32, 64, 128}
	fmt.Println(h)
	h = append(h, q...)
	fmt.Println(h)

	a := make([]int, 10, 100)
	fmt.Println(a, cap(a), len(a), &a)

}
