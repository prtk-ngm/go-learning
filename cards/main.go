package main

func main() {

	cards := newDeck()
	cards.saveToFile("my_cards")
	cards_file := newDeckFromFile("my_cards")
	cards_file.shuffle()
	cards_file.print()

}
