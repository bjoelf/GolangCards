package main

func main() {
	cards := newDeck()

	cards.saveToFile("my_cards")

	newCards := newDeckFromFile("my_cards")
	newCards.print()

	newCards.shuffle()
	newCards.print()

	//fmt.Println(cards.toString())
	//hand, remainingCards := deal(cards, 5)

	//hand.print()
	//remainingCards.print()

	//greeting := "Hi there"
	//fmt.Println([]byte(greeting))
}

//https://golang.org/pkg/io/ioutil/#WriteFile
