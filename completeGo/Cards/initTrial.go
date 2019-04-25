package main

import "fmt"

func trial() {

	//var card string = "Ace of Spades"
	card := newCard()
	fmt.Println(card)
	fmt.Printf("%T\n", card)

	cards := []string{newCard(), "Ace of Diamonds"}
	cards = append(cards, "Six of Spades")
	fmt.Println(cards)

	for i := range cards {
		fmt.Println(cards[i])
	}

}

func newCard() string {
	return "Five of Diamonds"
}
