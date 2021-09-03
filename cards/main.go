package main

func main() {
	/*cards := newDeck()

	hand, remainDeck := deal(cards, 5)

	hand.print()
	remainDeck.print() */

	/*cards := newDeck()
	cards.saveToFile("my_cards")*/

	/*cards := newDeckFromFile("my_card")
	cards.print() */
	cards := newDeck()
	cards.shuffle()
	cards.print()
}
