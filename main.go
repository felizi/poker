package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

/*
https://en.wikipedia.org/wiki/Playing_cards_in_Unicode
https://www.pokerstars.com/br/poker/games/rules/?no_redirect=1
https://br.888poker.com/how-to-play-poker/
https://pt.wikipedia.org/wiki/P%C3%B4quer
https://www.pokerstars.com/br/poker/games/rules/hand-rankings/
https://www.clubedopoker.com/o-valor-das-cartas-do-baralho-de-poker/
*/

const (
	spade   = "♤"
	heart   = "♡"
	diamond = "♢"
	club    = "♧"
	two     = "2"
	three   = "3"
	four    = "4"
	five    = "5"
	six     = "6"
	seven   = "7"
	eight   = "8"
	nine    = "9"
	ten     = "10"
	jack    = "J"
	queen   = "Q"
	king    = "K"
	aces    = "A"
)

func main() {
	var cards []Card
	sequence := []string{two, three, four, five, six, seven, eight, nine, ten, jack, queen, king, aces}
	suits := []string{spade, heart, diamond, club}
	for x := 0; x < len(sequence); x++ {
		for z := 0; z < len(suits); z++ {
			cards = append(cards, Card{ID: sequence[x], Suit: suits[z], Weight: x})
		}
	}

	player1 := hand(cards)
	fmt.Printf("player1: %v\n", player1)

	player2 := hand(remove(cards, player1[:]))
	fmt.Printf("player2: %v\n", player2)

	flop := flop(remove(cards, player2[:]))
	fmt.Printf("flop: %v\n", flop)

	turn := turn(remove(cards, flop[:]))
	fmt.Printf("turn: %v\n", turn)

	river := river(remove(cards, turn[:]))
	fmt.Printf("river: %v\n", river)

	fmt.Printf("community cards: %v\n", []Card{flop[0], flop[1], flop[2], turn[0], river[0]})

	weight, t := twoPair(player1, flop, turn, river)
	fmt.Printf("player1 - two pair: \t %v %v\n", weight, t)
	weight, c := onePair(player1, flop, turn, river)
	fmt.Printf("player1 - one pair: \t %v %v\n", weight, c)
	fmt.Printf("player1 - high card: \t %v\n", highCard(player1, flop, turn, river))

	weight, t = twoPair(player2, flop, turn, river)
	fmt.Printf("player2 - two pair: \t %v %v\n", weight, t)
	weight, c = onePair(player2, flop, turn, river)
	fmt.Printf("player2 - one pair: \t %v %v\n", weight, c)
	fmt.Printf("player2 - high card: \t %v\n", highCard(player2, flop, turn, river))
}

func remove(items []Card, exclusions []Card) []Card {
	for i := 0; i < len(items); i++ {
		for x := 0; x < len(exclusions); x++ {
			if items[i] == exclusions[x] {
				return remove(removeIndex(items, i), exclusions)
			}
		}
	}
	return items
}

func removeIndex(s []Card, index int) []Card {
	return append(s[:index], s[index+1:]...)
}

func hand(cards []Card) [2]Card {
	x := random(cards)
	y := random(cards)
	return [2]Card{x, y}
}

func flop(cards []Card) [3]Card {
	x := random(cards)
	y := random(cards)
	z := random(cards)
	return [3]Card{x, y, z}
}

func turn(cards []Card) [1]Card {
	return [1]Card{random(cards)}
}

func river(cards []Card) [1]Card {
	return [1]Card{random(cards)}
}

func random(x []Card) Card {
	rand.Seed(time.Now().UnixNano())
	return x[rand.Intn(len(x))]
}

func twoPair(hand [2]Card, flop [3]Card, turn, river [1]Card) (int, *[2][2]Card) {
	m := groupByID(hand, flop, turn, river)
	var matches []string
	for k, v := range m {
		if len(v) == 2 {
			matches = append(matches, k)
		}
	}
	if len(matches) == 2 {
		var result [2][2]Card
		var weight int
		for i := 0; i < len(matches); i++ {
			for x := 0; x < len(m[matches[i]]); x++ {
				c := m[matches[i]][x]
				weight += c.Weight
				result[i][x] = c
			}
		}
		return weight, &result
	}

	return 0, nil
}

func onePair(hand [2]Card, flop [3]Card, turn, river [1]Card) (int, *[2]Card) {
	cards := concatenate(hand, flop, turn, river)
	var card *Card
	for i := 0; i < len(cards); i++ {
		if card != nil && card.ID == cards[i].ID {
			return card.Weight + cards[i].Weight, &[2]Card{*card, cards[i]}
		}
		card = &cards[i]
	}
	return 0, nil
}

func groupByID(hand [2]Card, flop [3]Card, turn, river [1]Card) map[string][]Card {
	cards := concatenate(hand, flop, turn, river)
	m := make(map[string][]Card)
	for i := 0; i < len(cards); i++ {
		if m[cards[i].ID] == nil {
			m[cards[i].ID] = []Card{cards[i]}
		} else {
			m[cards[i].ID] = append(m[cards[i].ID], cards[i])
		}
	}
	return m
}

func highCard(hand [2]Card, flop [3]Card, turn, river [1]Card) int {
	cards := concatenate(hand, flop, turn, river)
	var sum int
	for i := 0; i < 5; i++ {
		sum += cards[i].Weight
	}
	return sum
}

func concatenate(hand [2]Card, flop [3]Card, turn, river [1]Card) [7]Card {
	cards := [7]Card{hand[0], hand[1], flop[0], flop[1], flop[2], turn[0], river[0]}
	sort.Slice(cards[:], func(i, j int) bool {
		return cards[i].Weight < cards[j].Weight
	})
	return cards
}

type Card struct {
	ID     string
	Suit   string
	Weight int
}
