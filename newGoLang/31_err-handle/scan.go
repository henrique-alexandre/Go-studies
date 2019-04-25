package main

import "fmt"

func main() {

	var a1, a2, a3 string

	fmt.Println("Name: ")
	_, err := fmt.Scan(&a1)
	if err != nil {
		panic(err)
	}

	fmt.Println("Favorite food: ")
	_, err = fmt.Scan(&a2)
	if err != nil {
		panic(err)
	}

	fmt.Println("Favorite sport: ")
	_, err = fmt.Scan(&a3)
	if err != nil {
		panic(err)
	}

	fmt.Println(a1, a2, a3)

}
