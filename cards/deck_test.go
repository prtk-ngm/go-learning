package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 16 {
		t.Errorf("Excpected deck length of 2000 but got %v", len(d))
	}

	if d[0] != "Ace of suit Spades" {
		t.Errorf("Excpected first card Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "Four of suit Clubs" {
		t.Errorf("Excpected last card Four of suit Clubs, but got %v", d[0])
	}

}

func TestSaveDeckToFileAndNewDeckFromFile(t *testing.T) {

	os.Remove("_deckteststring")
	d := newDeck()
	d.saveToFile("_deckteststring")
	loadedDeck := newDeckFromFile("_deckteststring")

	if d.toString() != loadedDeck.toString() {
		t.Errorf("Excpected deck string %v , but got  %v", d.toString(), loadedDeck.toString())
	}

}
