package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf(
			"Expected deck length of 20, but got %v",
			len(d),
		)
	}

	if d[0] != "Spades of Ace" {
		t.Errorf("Expected first card to be ace of spades but got %s", d[0])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	// Arrange
	filename := "_decktesting"
	os.Remove(filename)

	// Act
	deck := newDeck()
	deck.saveToFile(filename)

	// Assert
	loadedDeck := newDeckFromFile(filename)

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in deck, got %d", len(loadedDeck))
	}

	os.Remove(filename)
}
