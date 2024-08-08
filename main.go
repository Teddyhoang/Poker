package main

import "fmt"

func main() {
	decks := newDeck()
	// decks.print()
	// cards.printCard(3)
	var numPlayer int
	fmt.Print("Enter the number of player: ")
	fmt.Scanf("%d", &numPlayer)
	players := createPlayer(numPlayer)

	deal(&decks, players)
	// Print out the player deck
	for _, player := range players {
		fmt.Printf("ID: %s, Deck: %v\n", player.id, player.deck)
	}

	var playDeck deck
	state := "balanced"

	for i := range 3 {
		switch state {
		case "balanced":
			dealPlayDeck(&decks, &playDeck)
			i++
		}
	}

	fmt.Printf("final deck: %v\n", playDeck)
	decks.print()

	// tmp := players[0].deck
	// fmt.Println("Gia tri int cua card dau la: ", getCardValue(tmp[0]))
}
