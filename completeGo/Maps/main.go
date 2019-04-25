package main

import "fmt"

func main() {

	//var colors map[string]string //another way to declare a map
	colors := make(map[int]string) // equivalent way to create a map

	colours := map[string]string{

		"red":   "#ff0000",
		"green": "#00ff00",
		"white": "#ffffff",
	}

	colors[1] = "#ffffff"

	fmt.Println(colours)
	fmt.Println(colors)

	//delete(colours, "green")

	fmt.Println(colours)

	printMap(colours)

}

func printMap(m map[string]string) {
	for i, v := range m {
		fmt.Println(i, v)
	}
}
