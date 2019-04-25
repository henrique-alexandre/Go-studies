// File name: ...\s04\01_arrays_int\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)


//An array is a sequence of elements of a single type. The length is never negative

package main

import (
	"fmt"
)

func main() {
	var nums [3]int

	var sum1 int
	var sum2 int

	fmt.Printf("nums=%v type=%T len=%d\n", nums, nums, len(nums))

	for i := range nums {
		sum1 += i //increments the index
		sum2 += nums[i] //increment each position inside the array
	}

	fmt.Println(sum1, sum2)
}
