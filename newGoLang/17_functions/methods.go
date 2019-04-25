package main

import "fmt"

//a value can be of more than one type
type human interface {
	speak()
}

type person struct {
	fName string
	lName string
	age   int
}

type professional struct {
	person
	cert bool
}

func main() {

	o := foo()
	s := fmt.Sprintf("%T\n", o)
	fmt.Print(s)
	//u := o()
	//fmt.Println(u)

	//func expression = we assign a function to a variable
	f := func() { fmt.Println("My first func expression") }
	f()

	//anonymous func
	func(x int) {
		fmt.Println("Hi", x)
	}(50)

	p1 := professional{
		person{
			"Henrique",
			"Alexandre",
			32,
		},
		true,
	}

	p2 := person{
		"Todd",
		"Mcleodd",
		45,
	}

	fmt.Println(p1.fName)
	p1.speak()
	bar(p1)
	bar(p2)
}

func (p person) speak() {
	fmt.Println("I am ", p.fName, p.lName)
}

func bar(h human) {

	fmt.Println("I'm a bloddy human!")
}

func foo() func() int {
	return func() int {
		return 500
	}
}
