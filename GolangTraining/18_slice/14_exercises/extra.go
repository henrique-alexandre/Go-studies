package main

import "fmt"

func main() {

	customerNumber := make([]int, 0)

	customerNumber[0] = 7
	customerNumber[1] = 10
	customerNumber[2] = 15

	fmt.Println(customerNumber[0])
	fmt.Println(customerNumber[1])
	fmt.Println(customerNumber[2])

	greeting := make([]string, 3, 5)

	greeting[0] = "Olá"
	greeting[1] = "Manishma"
	greeting[2] = "Hi"

	fmt.Println(greeting[1])

}
