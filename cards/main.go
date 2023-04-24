package main

func main() {
	// var card string = "Ace of Spades"
	// cards := newDeck()
	// cards.saveToFile("my_cards")
	// cards := newDeckFromFile("my_cards1")
	// cards.print()

	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }
	// cards.print()
	// fmt.Println(cards)
	// hand, remainingCards := deal(cards, 5)
	// hand.print()
	// remainingCards.print()

	cards := newDeck()
	cards.shuffle()
	cards.print()
}
