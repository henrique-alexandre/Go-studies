package main

import "fmt"

//control flow: how the computer reads the code. Sequential reading.

func main() {

	//for structure: init -- condition -- post

	fmt.Println("Henrique")
	fmt.Println("Alexandre")

	for i := 1; i <= 10; i++ {
		fmt.Printf("The outer loop is: %d\n", i)
		for j := 1; j < 4; j++ {
			fmt.Printf("\t\t The inner loop is: %d.%d\n", i, j)
		}
	}

	for i := 1; i <= 50; i++ {

		if i%2 == 0 {
			fmt.Printf("%v ", i)
		}
	}

	fmt.Println()

	for i := 65; i <= 122; i++ {

		fmt.Printf("%#U\n", i)

	}

	for i := 65; i <= 122; i++ {

		fmt.Println(string(i))
	}

}
