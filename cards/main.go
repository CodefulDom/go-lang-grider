package main

func main() {
	cards := newDeck()

	hand, reaminaingCards := deal(cards, 5)

	hand.print()
	reaminaingCards.print()
}
