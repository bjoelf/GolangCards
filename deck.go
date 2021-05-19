package main

//https://golang.org/pkg/fmt/
import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
)

//Create a new type of deck, wich is a slice of strings.
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuites := []string{"Spades", "Dimonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuites {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)

	if err != nil {
		//Option #1 - log the error and return a call to newDeck()
		//Option #2 - log the error and entirely quit the program
		fmt.Println("Error: ", err)
		//exit with package os
		os.Exit(1)
	}
	//biteslice into string, cast
	//cardSuites := []string{"Spades", "Dimonds", "Hearts", "Clubs"}

	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffle() {
	for i := range d {
		//https://golang.org/pkg/math/rand/#Intn
		newPosition := rand.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
