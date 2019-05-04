package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []card

type card struct {
	value string
	class string
}

func newDeck() deck {
	cardClasses := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	cards := deck{}

	for _, class := range cardClasses {
		for _, value := range cardValues {
			cards = append(cards, card{class: class, value: value})
		}
	}
	return cards
}

func (d deck) print() { //d actual copy of the deck we'r working with like this or self
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		swapPosition := r.Intn(len(d) - 1) //rand.Intn always use same exact seed for rand num generation. So we randomized our own seed source with current time.

		d[i], d[swapPosition] = d[swapPosition], d[i]
	}

}

func deal(d deck, handSize int) (deck, deck) { //create new hand of given size.
	return d[:handSize], d[handSize:]
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	strSlice := strings.Split(string(bs), ",")
	cards := []card{}
	for _, strValue := range strSlice {
		cardAttr := strings.Split(strValue, " of ")
		newCard := card{cardAttr[0], cardAttr[1]}
		cards = append(cards, newCard)
	}

	return deck(cards)
}

func (d deck) toString() string {
	cards := []card(d)

	strArrOfCards := []string{}

	for _, card := range cards {
		strArrOfCards = append(strArrOfCards, card.value+" of "+card.class)
	}
	return strings.Join(strArrOfCards, ",")
}

func (c card) toString() string {
	return c.value + " of " + c.class
}
