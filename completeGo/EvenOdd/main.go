package main

import "fmt"

func main() {

	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for v := range s {
		if v%2 == 0 {
			fmt.Printf("The number %v is even\n", v)
		} else {
			fmt.Printf("The number %v is odd\n", v)
		}
	}

}
