package main

import "fmt"

func main() {

	defer fmt.Println(fo(1, 2, 3, 4, 5, 6, 7, 8, 9))
	s := []int{90, 80, 70, 60, 50, 40, 30, 20, 10}
	fmt.Println(ba(s))

}

func fo(n ...int) int {

	//i := []int{}
	//i = n

	total := 0
	for _, v := range n {
		total += v

	}
	return total

}

func ba(ai []int) int {

	total := 0
	for _, v := range ai {
		total += v

	}
	return total

}
