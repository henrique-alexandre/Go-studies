// File name: ...\s04\07_slices_len_cap\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

//A slice is a descriptor for a contiguous segment of an underlying array
//and provides access to a numbered sequence of elements from that array.
//A slice type denotes the set of all slices of arrays of its element type.
//The value of an uninitialized slice is nil.

package main

import (
	"fmt"
)



func main() {
	days := [...]string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}

	fmt.Printf("days=%v\n\n", days) //	or days[:]

	weekdays := days[:5] //or days[0:5] including 0, excluding 5; called 'slicing'
	fmt.Printf("weekdays=%v\n", weekdays)
	fmt.Printf("len(weekdays)=%d\n", len(weekdays))
	fmt.Printf("cap(weekdays)=%d\n", cap(weekdays))
	fmt.Printf("weekdays[0])=%s\n", weekdays[0])
	fmt.Printf("weekdays[4])=%s\n\n", weekdays[4])
	// fmt.Println(weekdays[5]) //panic: runtime error: index out of range

	weekend := days[5:] //or days[5:7]
	fmt.Printf("weekend=%v\n", weekend)
	fmt.Printf("len(weekend)=%d\n", len(weekend))
	fmt.Printf("cap(weekend)=%d\n", cap(weekend))
	fmt.Printf("weekend[0])=%s\n\n", weekend[0])
	// fmt.Println(weekend[2]) //panic: runtime error: index out of range

	retired := days[:0]
	fmt.Printf("retired=%v\n", retired)
	fmt.Printf("len(retired)=%d\n", len(retired))
	fmt.Printf("cap(retired)=%d\n\n", cap(retired))
	// fmt.Println(retired[0])	//panic: runtime error: index out of range

	fmt.Printf("my vacation=%v\n", days[3:5])
}
