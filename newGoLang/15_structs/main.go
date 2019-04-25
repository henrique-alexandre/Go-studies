package main

import "fmt"

//Struct is a data structure which allows us to compose together values of different types.
//It aggregates together values of different types

type person struct {
	firstName string
	lastName  string
	age       int
}

type secretAgent struct {
	person
	ltk bool
}

func main() {

	p1 := person{
		firstName: "Henrique",
		lastName:  "Alexandre",
		age:       32,
	}

	p2 := struct {
		firstName string
		lastName  string
		age       int
	}{
		firstName: "Todd",
		lastName:  "Mcleodd",
		age:       40,
	}

	sa1 := secretAgent{

		person: p1,
		ltk:    true,
	}

	fmt.Println(p1.firstName, p1.lastName, p1.age)
	fmt.Println(sa1.ltk)
	fmt.Println(p2)

}
