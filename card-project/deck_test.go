package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 52 {
		t.Errorf("Expected deck lenth of 52, got %d", len(d))
	}

	if d[0].toString() != "Ace of Spades" {
		t.Errorf("Expected first card to be Ace of Spades but found %s", d[0])
	}

	if d[len(d)-1].toString() != "King of Clubs" {
		t.Errorf("Expected last card to be Clubs of Kings but found %s", d[len(d)-1])
	}
}

func TestSaveToFileandNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	d := newDeck()
	d.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 52 {
		t.Errorf("Expected 52 cards in loaded deck, got %d", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
