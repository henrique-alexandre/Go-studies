package main

import "fmt"

func main() {

	a := "Hello"

	i, _ := fmt.Println("Ahava", a)
	fmt.Println(i)

	foo()

	for i := 0; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Print(i, " ")
		}

	}

	fmt.Println()

	bar()

}

func foo() {
	fmt.Println("I'm in foo")
}

func bar() {

	fmt.Println("we exited")
}

//types of control flow: (1) sequence; (2) loop iterative; (3)conditional
