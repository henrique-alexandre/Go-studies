package main

import (
	"testing"
)

func TestNewDeck(t *testing.T) {

	d := newDeck()

	if len(d) != 36 {
		t.Errorf("Expected deck lenght of 36, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected AoS but got %v ", d[0])
	}

	if d[len(d)-1] != "Nine of Club" {
		t.Errorf("Expexted NoC, but got %v ", d[len(d)-1])
	}

}
