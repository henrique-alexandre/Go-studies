package main

//the command 'import' creates a link that imports resources from other packages.
import "fmt"

//package = a collection of common source files
//The name 'main' determines the package is executable. Any other name defines reusability, dependencies, not execution.
//func main is required for execution -- entry point.

type person struct {
	fName string
	lName string
}

func main() {

	fmt.Println("Henrique")

	mySlice := []string{"Be", "Kind"} //A slice is composed of a pointer to an array, size (length) and capacity values.

	fmt.Println(mySlice)

	updateSlice(mySlice)

	fmt.Println(mySlice)

	p1 := person{
		"Henrique",
		"Alexandre",
	}

	fmt.Println(p1)

	updatePerson(p1)

	fmt.Println(p1)

}

func updateSlice(s []string) {
	s[0] = "Bye"

}

func updatePerson(p person) {
	p.fName = "Pivotal"
}
