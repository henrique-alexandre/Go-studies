package main

import "fmt"

func main() {

	num := 52
	fmt.Printf("%b \n", num)

	mySlice := make([]int, 0)
	fmt.Println("---------------")
	fmt.Println(mySlice)
	fmt.Println(len(mySlice))
	fmt.Println(cap(mySlice))
	fmt.Println("---------------")

	for i := 0; i < 81; i++ {
		mySlice = append(mySlice, i)
		fmt.Println("Length: ", len(mySlice), " Capacity: ", cap(mySlice), " Value: ", mySlice[i])

		S2 := []int{10, 15, 20, 25, 30, 35, 40, 45, 50}

		for i, n := range S2 {
			fmt.Println(i, " - ", n)

		}

	}
}
