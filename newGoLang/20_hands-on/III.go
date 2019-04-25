package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println(p.first)
	fmt.Println(p.age)

}

func main() {

	p1 := person{
		"Henrique",
		"Alexandre",
		32,
	}

	p1.speak()
}
