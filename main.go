package main

func main() {
	cards := deck{"Six of Spades", newCard()}
	cards = append(cards, "Ace of Spades")

	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
