package main

import "fmt"

func main() {

	a := []int{10, 20, 30, 40, 50}
	a = append(a, 52, 53, 54, 55)
	fmt.Println(a)
	b := []int{56, 57, 58, 59, 60}
	a = append(a, b...)
	fmt.Println(a)

	for index, value := range a {
		fmt.Printf("The position %v is holding the value %v\n.", index, value)

	}
	fmt.Printf("%T\n", a)

	ss := map[string][]string{
		"James Bond": []string{"shaken", "martini", "women"},
		"Miss":       []string{"literature", "computer science"},
	}

	ss["Ian"] = []string{"steaks", "espionage"}

	for k, v := range ss {
		fmt.Println("This is the record", k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}

		delete(ss, "Ian")
	}

}
