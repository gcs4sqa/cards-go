package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 48 {
		t.Errorf("Expected deck length of 48, but got %v", len(d))
	}

	if d[0] != "Ace of Clubs" {
		t.Errorf("Expected first card of Ace of Clubs but got %v", d[0])
	}

	if d[len(d)-1] != "King of Diamonds" {
		t.Errorf("Expected first card of King of Diamonds but got %v", d[len(d)-1])

	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 48 {
		t.Errorf("Expected length of deck was 48, but got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
