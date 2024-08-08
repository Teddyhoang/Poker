package main

import (
	"fmt"
	"os"
)

type player struct {
	id int
	deck
}

// Declare player
func createPlayer(numPlayer int) []player {
	if numPlayer >= 10 {
		fmt.Println("Exceeding the allowed number of players")
		os.Exit(1)
	}

	players := make([]player, numPlayer)

	for i := range numPlayer {
		players[i] = player{id: i, deck: []card{}}
	}

	return players
}

// func bet(betValue int) {

// }
