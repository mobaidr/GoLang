package main

import "testing"

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
