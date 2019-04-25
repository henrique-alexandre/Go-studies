package main

import "fmt"

type person struct {
	fName string
	lName string
	age   int
}

func (p *person) speak() {
	fmt.Println("Olá")
}

type human interface {
	speak()
}

func main() {

	p1 := person{
		"Henrique",
		"Alexandre",
		32,
	}

	saySomething(&p1)
	//saySomething(p1)

	p1.speak()

}

func saySomething(h human) {
	h.speak()
}
