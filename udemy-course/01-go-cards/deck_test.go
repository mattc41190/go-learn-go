package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 52 {
		t.Errorf("Expected length of 52, but got %v", len(d))
	}
}

func TestFirstCard(t *testing.T) {
	d := newDeck()
	if d[0] != "Ace of Hearts" {
		t.Errorf("Expected: Ace of Hearts but go %v", d[0])
	}
}

func TestLastCard(t *testing.T) {
	d := newDeck()
	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected: King of Clubs but go %v", d[len(d)-1])
	}
}

func TestSaveDeck(t *testing.T) {
	os.Remove("_decktesting.txt")
	d := newDeck()
	d.saveToFile("_decktesting.txt")
	d2 := newDeckFromFile("_decktesting.txt")
	if len(d2) != 52 {
		t.Errorf("Expected length of 52, but got %v", len(d2))
	}
	os.Remove("_decktesting.txt")
}
