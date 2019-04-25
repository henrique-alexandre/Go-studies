package main

import "fmt"

type vehicle struct {
	doors  int
	colour string
}

type truck struct {
	fourWheel bool
	vehicle
}

type sedan struct {
	luxury bool
	vehicle
}

func main() {
	t := truck{

		false,
		vehicle{

			4,
			"Black",
		},
	}

	s := sedan{
		true,
		vehicle{
			5,
			"White",
		},
	}

	fmt.Println(t)
	fmt.Println(s)

	o := struct {
		name      string
		friends   map[string]int
		favDrinks []string
	}{
		name: "David",
		friends: map[string]int{
			"Richard":   555,
			"Sedutor":   888,
			"Malfadado": 111,
		},
		favDrinks: []string{"cachaça", "pinga", "velho barreiro"},
	}

	fmt.Println(o.name)
	fmt.Println(o.friends)
	fmt.Println(o.favDrinks)

}
