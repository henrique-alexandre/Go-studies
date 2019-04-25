package main

import "fmt"

func main() {
	x := 13 % 3
	fmt.Print(x)
	if x == 1 {
		fmt.Print(" is Odd")
	} else {
		fmt.Print(" is Even")
	}

	fmt.Println()

	for i := 1; i <= 70; i++ {
		if i%2 == 1 {
			fmt.Printf("%v is Odd", i)
			fmt.Println()
		} else {
			fmt.Printf("%v is Even", i)
			fmt.Println()
		}
	}
}
