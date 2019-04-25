package main

import "fmt"

type person struct {
	fName  string
	lName  string
	favIcc []string
}

func main() {

	p1 := person{

		fName:  "Henrique",
		lName:  "Alexandre",
		favIcc: []string{"mint", "pistache", "chocolate"},
	}

	p2 := person{

		"Todd",
		"Mcleodd",
		[]string{"strawberry", "gum", "kinder"},
	}

	for _, v := range p1.favIcc {
		fmt.Println(v)

	}
	for _, v := range p2.favIcc {
		fmt.Println(v)

	}

	m := map[string]person{

		p1.lName: p1,
		p2.lName: p2,
	}

	for k, v := range m {

		fmt.Println(k)
		fmt.Println(v.fName)
		fmt.Println(v.lName)
		for i, val := range v.favIcc {
			fmt.Println(i, val)
		}
	}
}
