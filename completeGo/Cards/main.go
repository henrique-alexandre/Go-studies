package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func main() {

	//cards := deck{"Ace of Spades", "Five of diamonds"}
	//
	//fmt.Println(cards)

	//h := "Henrique"
	//fmt.Println([]byte(h))

	//cards := newDeck()
	//cards.saveToFile("teste")
	//cards.print()
	//hand, remainingCards := deal(cards, 5)

	//hand.print()
	//remainingCards.print()

	//fmt.Println(cards.toSB())

	cards := newDeckFromFile("teste")
	cards.shuffle()
	cards.print()

}

func (d deck) print() {
	for i, v := range d {
		fmt.Printf("%v %v \n ", i, v)
	}
}

func newDeck() deck {

	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Club"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine"}

	for _, value := range cardValues {
		for _, suit := range cardSuits {
			cards = append(cards, value+" of "+suit)

		}
	}
	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toSB() []byte {
	ns := strings.Join([]string(d), ",")
	y := []byte(ns)
	return y
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, d.toSB(), 777)

}

func newDeckFromFile(filename string) deck {

	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	d := deck(s)

	return d
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		//newPosition := rand.Intn(len(d) - 1) //run without NewSource
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
