package main

import "fmt"

//functions must be modularised
// Functions signature: func (r receiver) identifier(parameters) (return (s)) {code}
//we define our func with parameters, but we call it passing arguments in ()
//everything in Go is PASSED BY VALUE. What you see is what you get

func main() {

	//unfurling a Slice. '...T is equivalent to []T'
	xi := []int{2, 3, 4, 5, 6, 7, 8, 9}
	sum(xi...)

	s2 := odd(sum, xi...)
	fmt.Println(s2)

}

//Variadic means 0 or more
func sum(x ...int) int {

	defer fmt.Printf("%T\n", x)

	sum := 0
	for i, v := range x {
		sum += v
		fmt.Println("The position ", i, "is ", sum)
	}

	fmt.Println("The total sum is ", sum)
	fmt.Println(x)

	return sum
}

//example of callback
func odd(f func(xi ...int) int, vi ...int) int {
	var oi []int
	for _, v := range vi {
		if v%2 != 0 {
			oi = append(oi, v)
		}
	}
	return f(oi...)
}
