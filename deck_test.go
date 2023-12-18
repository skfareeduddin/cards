package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if (len(d) != 52) {
		t.Errorf("Expected length of 52, but got %v", len(d))
	}

	if (d[0] != "Ace of Spades") {
		t.Errorf("Expected card - Ace of Spades, but found card - %v", d[0])
	}

	if (d[len(d) - 1] != "King of Clubs") {
		t.Errorf("Expected card - King of Clubs, but found card - %v", d[len(d) - 1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if (len(loadedDeck) != 52) {
		t.Errorf("Expected 52 cards, got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}