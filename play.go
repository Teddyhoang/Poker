package main

import (
	"sort"
)

type valueDeck struct {
	playerId   int
	valueCards card
	valueState int
}

func dealPlayDeck(d *deck, pd *deck) {
	(*d) = (*d)[1:]
	if len(*pd) == 0 {
		for i := range 3 {
			(*pd) = append((*pd), (*d)[0])
			(*d) = (*d)[1:]
			i++
		}
	} else {
		(*pd) = append((*pd), (*d)[0])
		(*d) = (*d)[1:]
	}
}

// func getWinner(pd *deck, p []player) player {
// 	winner := p[0]
// 	if len(*pd) == 5 {
// 		winner = bestDeck((*pd), p)
// 	}
// 	return winner
// }

// func bestDeck(pd *deck, p []player) player {

// }

func seven(pd deck, pl player) (id int, strength int) {
	var tmp7Cards []card
	tmp7Cards = make([]card, 7)

	var tmp5Cards []card
	tmp5Cards = make([]card, 5)

	var best []card
	best = make([]card, 5)

	id = 1
	strength = 0
	tmp7Cards = append(pd, pl.deck...)

	for i, item := range t7c5 {
		tmp5Cards[0] = tmp7Cards[item[0]]
		tmp5Cards[1] = tmp7Cards[item[1]]
		tmp5Cards[2] = tmp7Cards[item[2]]
		tmp5Cards[3] = tmp7Cards[item[3]]
		tmp5Cards[4] = tmp7Cards[item[4]]

		tmpstrength := evaluated(tmp5Cards)
		if tmpstrength > strength {
			strength = tmpstrength
			id = i + 1
			best = tmp5Cards
		}
		if tmpstrength == strength {
			compareHand(tmpstrength, best)
		}
	}

	return id, strength

}

// func evaluated(d deck) (s int) {

// }

func hasStraight(decks deck) bool {
	values := make([]int, 5)
	for i, card := range decks {
		values[i] = getCardValue(card)
	}

	sort.Ints(values)

	isAceLowStraight := func(value []int) bool {
		aceLow := []int{2, 3, 4, 5, 14}
		count := 0
		for _, v := range value {
			if count < 5 && v == aceLow[count] {
				count++
			}
		}

		return count == 5
	}

	isRegularStraight := func(value []int) bool {
		count := 0
		for i := 1; i < len(value); i++ {
			if value[i] == value[i-1]+1 {
				count++
			}
		}

		return count == 5
	}

	return (isRegularStraight(values) || isAceLowStraight(values))
}

func hasFlush(decks deck) bool {
	value := make([]int, 5)
	for i, card := range decks {
		value[i] = getSuitValue(card)
	}

	count := 0

	for i := 1; i < len(value); i++ {
		if value[i] == value[i-1] {
			count++
		}
	}

	return count == 5
}

func getNumUniqueValue(cards []card) map[int]int {
	uniqueMap := make(map[int]int)
	values := make([]int, 5)
	for i, card := range cards {
		values[i] = getCardValue(card)
	}

	for _, value := range values {
		uniqueMap[value]++
	}

	return uniqueMap
}

func findCombination(m map[int]int) int {
	count4 := 0
	count3 := 0
	count2 := 0

	for _, value := range m {
		switch value {
		case 4:
			count4++
		case 3:
			count3++
		case 2:
			count2++
		}
	}

	if count4 > 0 {
		return 3
	}
	if count3 > 0 && count2 > 0 {
		return 4
	}
	if count3 > 0 && count2 == 0 {
		return 7
	}
	if count2 > 1 {
		return 8
	}
	if count2 == 1 {
		return 9
	}
	return 0
}

var t7c5 = [21][7]uint8{
	{0, 1, 2, 3, 4, 5, 6},
	{0, 1, 2, 3, 5, 4, 6},
	{0, 1, 2, 3, 6, 4, 5},
	{0, 1, 2, 4, 5, 3, 6},
	{0, 1, 2, 4, 6, 3, 5},
	{0, 1, 2, 5, 6, 3, 4},
	{0, 1, 3, 4, 5, 2, 6},
	{0, 1, 3, 4, 6, 2, 5},
	{0, 1, 3, 5, 6, 2, 4},
	{0, 1, 4, 5, 6, 2, 3},
	{0, 2, 3, 4, 5, 1, 6},
	{0, 2, 3, 4, 6, 1, 5},
	{0, 2, 3, 5, 6, 1, 4},
	{0, 2, 4, 5, 6, 1, 3},
	{0, 3, 4, 5, 6, 1, 2},
	{1, 2, 3, 4, 5, 0, 6},
	{1, 2, 3, 4, 6, 0, 5},
	{1, 2, 3, 5, 6, 0, 4},
	{1, 2, 4, 5, 6, 0, 3},
	{1, 3, 4, 5, 6, 0, 2},
	{2, 3, 4, 5, 6, 0, 1},
}
