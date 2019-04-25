package main

import "fmt"

func main() {

	fmt.Println(mySum(5 + 5))
	fmt.Println(mySum(50 + 100))

}

func mySum(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum + 1
}
