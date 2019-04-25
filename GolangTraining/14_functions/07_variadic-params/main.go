package main

import "fmt"

func main() {
	n := average(40, 50, 80, 10, 40, 50)
	fmt.Println(n)
}

func average(sf ...float64) float64 {
	fmt.Println(sf)
	fmt.Printf("The type is %T \n", sf)
	var total float64
	//fmt.Println("the initial total is: ", total)
	for _, v := range sf {
		total += v
	}
	return total / float64(len(sf))
}
