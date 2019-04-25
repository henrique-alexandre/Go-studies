package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type family struct {
	person
	part_of_Family bool
}

type person struct {
	name    string
	age     int
	height  float32
	address string
}

func (p person) addsurname(s string) string {
	p.name = p.name + s
	return p.name
}

type stroll interface {
	goPark() bool
}

func (f family) goPark() bool {
	return true
}

func (p person) goPark() bool {
	return true
}

func add(x, z int) int {
	return x + z
}

func main() {

	p1 := person{
		"Henrique",
		32,
		1.75,
		"Dublin",
	}

	p1.addsurname("Alexandre")
	//fmt.Println(p1)

	f1 := family{
		p1,
		true,
	}

	p1.goPark()
	f1.goPark()

}

func extra() {
	xi := []int{2, 4, 6, 8, 10}
	fmt.Println(xi)

	m := map[string]int{

		"Henrique": 32,
		"Cilene":   28,
	}
	fmt.Println(m)

	num := 9
	fmt.Printf("The type is %T \n", num)

	fmt.Println(time.Now(), math.Pi)
	fmt.Println("My favourite number is: ", rand.Intn(20))
	fmt.Printf("You have: %g \n", math.Sqrt(16))
	fmt.Println(add(10, 20))

	//for i:=60; i<=121; i++ {
	//	fmt.Printf("%v \t %b \t %x \t %q \n", i, i, i, i)
	//}

}
