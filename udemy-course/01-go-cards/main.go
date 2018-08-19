package main

func main() {
	cards := newDeckFromFile("my_cards.txt")
	// hand, remainingCards := deal(cards, 5)
	cards.shuffle()
	cards.print()
}
