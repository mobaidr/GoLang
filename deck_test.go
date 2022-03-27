package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 20, but got %d", len(d))
	}

	expectedStr := "Ace of Spades"
	actualStr := d[0]
	if d[0] != expectedStr {
		t.Errorf("Expected %s, but got %s", expectedStr, actualStr)
	}

	expectedStr = "Four of Clubs"
	actualStr = d[len(d)-1]
	if actualStr != expectedStr {
		t.Errorf("Expected %s, but got %s", expectedStr, actualStr)
	}
}

func TestSaveToDeckAndNewDeckTestFromFile(t *testing.T) {
	filename := "_ desktesting"

	os.Remove(filename)

	deck := newDeck()

	deck.saveToFile(filename)

	loadedDeck := newDeckFrom(filename)

	if len(loadedDeck) != 16 {
		t.Errorf("Expected deck length of 20, but got %d", len(loadedDeck))
	}

	expectedStr := deck[0]
	actualStr := loadedDeck[0]
	if loadedDeck[0] != expectedStr {
		t.Errorf("Expected %s, but got %s", expectedStr, actualStr)
	}

	expectedStr = deck[len(deck)-1]
	actualStr = loadedDeck[len(loadedDeck)-1]
	if actualStr != expectedStr {
		t.Errorf("Expected %s, but got %s", expectedStr, actualStr)
	}

	os.Remove(filename)
}
