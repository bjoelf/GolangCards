package main

func main() {

	cards := deck{"Ace of Diamonds", newCard()}
	cards = append(cards, "another slice")

	cards.print()
}

func newCard() string {
	return "Five of Dimonds"
}

func timesTwoNumber(number int) int {
	return number * 2
}
