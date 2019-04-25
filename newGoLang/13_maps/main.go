package main

import "fmt"

func main() {

	//Maps are unordered lists. Complex object composed of keys and values.
	//A map is a built-in data structure that associates values using keys and values.

	m := map[string]int{
		"Henrique": 32,
		"Todd":     40,
	}
	fmt.Println(m)
	fmt.Println(m["Henrique"], m["Todd"])

	v, ok := m["Henrique"]
	fmt.Println(v, ok)

	if v, ok := m["Todd"]; ok {
		fmt.Println(v)
	} else {
		fmt.Print("user not found")
	}

	for k, v := range m {
		fmt.Println(k, v)

	}

	delete(m, "Todd")
	fmt.Println(m)

}
