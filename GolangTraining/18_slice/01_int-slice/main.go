package main

import "fmt"

func main() {

	mySlice := []int{1, 3, 5, 7, 9, 11}
	fmt.Printf("%T\n", mySlice)
	fmt.Println(mySlice)
	fmt.Println(mySlice[2:4])
	fmt.Println("myString"[2]) //the index refers to the String written just before it!

}

//Slice is a data structure we use for storing a LIST of values of a certain type.
//Slices are mode of a pointer to an underlying array, a length and capacity.
//Slice is a REFERENCE type. A slice is built on top of an array.
//The length is a number of elements the slice currently holds. We can access it using index positions.
//Capacity is the number of elements available in the underlying array.
//If we don't initialize a slice there is no underlying array.
//We can delete items from the slice using the append method. eg.: mySlice = append(mySlice[:2], mySlice[3:])
