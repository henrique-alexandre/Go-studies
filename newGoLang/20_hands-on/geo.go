package main

import (
	"fmt"
	"math"
)

type square struct {
	l float32
}

type circle struct {
	r float32
}

func (s square) area() float32 {
	return s.l * s.l

}
func (c circle) area() float32 {
	return math.Pi * c.r * c.r
}

type shape interface {
	area() float32
}

func info(s shape) {
	fmt.Println(s.area())
}

func main() {

	aa := square{
		5}

	bb := circle{
		25}

	info(aa)
	info(bb)

}
