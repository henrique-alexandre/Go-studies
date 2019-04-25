package main

import "fmt"

// every value and function in Go has a type
//Interface type is the opposite of concrete type.

type spanishBot struct{}
type englishBot struct{}

type bot interface {
	getGreeting() string
}

func main() {

	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)

}

func (sb spanishBot) getGreeting() string {
	return "Ola"
}
func (eb englishBot) getGreeting() string {
	return "Hello"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())

}
