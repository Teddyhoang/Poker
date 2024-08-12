package main

import (
	"fmt"
	"math/rand"
	"time"
)

type card struct {
	value string
	suit  string
}

type deck []card

var cardValueMap = map[string]int{
	"Two":   2,
	"Three": 3,
	"Four":  4,
	"Five":  5,
	"Six":   6,
	"Seven": 7,
	"Eight": 8,
	"Nine":  9,
	"Ten":   10,
	"Jack":  11,
	"Queen": 12,
	"King":  13,
	"Ace":   14,
}

func getCardValue(c card) int {
	return cardValueMap[c.value]
}

var suitValueMap = map[string]int{
	"Spades":   3,
	"Hearts":   2,
	"Diamonds": 1,
	"Clubs":    0,
}

func getSuitValue(c card) int {
	return cardValueMap[c.suit]
}

// Create Deck
func newDeck() deck {
	desks := deck{}
	cardsSuit := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardsValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
	var tmp card

	for _, s := range cardsSuit {
		for _, v := range cardsValues {
			tmp.updateSuit(s)
			tmp.updateValue(v)
			desks = append(desks, tmp)
		}
	}

	desks.shuffle()

	return desks
}

func (d *card) updateSuit(newSuit string) {
	(*d).suit = newSuit
}

func (d *card) updateValue(newValue string) {
	(*d).value = newValue
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		randNum := r.Intn(len(d) - 1)
		d[i], d[randNum] = d[randNum], d[i]
	}
}

// Deal functions
func deal(d *deck, players []player) {
	for i := range 2 {
		for j := range len(players) {
			players[j].deck = append(players[j].deck, (*d)[0])
			(*d) = (*d)[1:]
		}
		i++
	}
}

// Print functions
func (d deck) print() {
	fmt.Printf("Number of cards left: %d\n", len(d))
	fmt.Println(d)
}

func (d deck) printCard(i int) {
	fmt.Println(d[i])
}
