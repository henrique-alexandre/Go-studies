package main

import "fmt"

type square struct {
	side float64
}

type triangle struct {
	base   float64
	height float64
}

type shape interface {
	getArea() float64
}

func main() {

	c1 := square{
		10,
	}

	t1 := triangle{
		10,
		10,
	}

	printArea(t1)
	printArea(c1)

}

func (s square) getArea() float64 {
	return s.side * s.side
}
func (t triangle) getArea() float64 {
	return 0.5 * t.height * t.base
}

func printArea(a shape) {
	fmt.Println(a.getArea())
}
