package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	//We MUST have all fields beginning with an uppercase letter, otherwise they won't be exported
	First string
	Last  string
	Age   int
}

//A String is a sequence of bytes

func main() {

	p1 := person{
		"Henrique",
		"Alexandre",
		32,
	}
	p2 := person{
		"Todd",
		"Mcleodd",
		42,
	}

	people := []person{p1, p2}
	people10 := []person{p1, p2}

	fmt.Println(people)

	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))

	err = json.Unmarshal(bs, &people10)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(people10)

}
