package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
	contact   contactInfo
}

type contactInfo struct {
	email string
	eir   string
}

func main() {
	p1 := person{
		firstName: "Henrique",
		lastName:  "Alexandre",
		age:       32,
		contact: contactInfo{
			email: "henriquealexandre@gmail.com",
			eir:   "D24-PF76",
		},
	}

	fmt.Printf("%v\n", p1)

	p1.lastName = "Cordeiro"

	p1.print()

	p1.updateName("Marcio")

	p1.print()
}

func (p person) print() {

	fmt.Printf("%+v\n", p)

}
func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
	fmt.Println(p.firstName)
	fmt.Println(&p.firstName)

}
