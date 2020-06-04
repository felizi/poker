package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

/*
https://en.wikipedia.org/wiki/Playing_cards_in_Unicode
https://www.pokerstars.com/br/poker/games/rules
https://br.888poker.com/how-to-play-poker/
https://pt.wikipedia.org/wiki/P%C3%B4quer
https://www.pokerstars.com/br/poker/games/rules/hand-rankings/
https://www.clubedopoker.com/o-valor-das-cartas-do-baralho-de-poker/
https://www.cardplayer.com/rules-of-poker/hand-rankings
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

// Hand of Poker
type Hand int

const (
	HighCard      Hand = iota
	OnePair            = iota
	TwoPair            = iota
	ThreeOfAKind       = iota
	Straight           = iota
	Flush              = iota
	FullHouse          = iota
	FourOfAKind        = iota
	StraightFlush      = iota
	RoyalFlush         = iota
)

func (d Hand) String() string {
	return [...]string{"High Card", "One Pair", "Two Pair", "Three of a Kind", "Straight", "Flush", "Full House", "Four of a Kind", "Straight Flush", "Royal Flush"}[d]
}

type Checker interface {
	execute(cards [7]Card) (Hand, int, *[5]Card)
}

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
	turn := turn(remove(cards, flop[:]))
	river := river(remove(cards, turn[:]))
	fmt.Printf("community cards: %v\n", []Card{flop[0], flop[1], flop[2], turn[0], river[0]})

	h, w, c := check(player1, flop, turn, river)
	fmt.Printf("Player1:\t[%v]\t[%v]\t%v\n", *h, *w, *c)
	h, w, c = check(player2, flop, turn, river)
	fmt.Printf("Player2:\t[%v]\t[%v]\t%v\n", *h, *w, *c)
}

func check(hand [2]Card, flop [3]Card, turn, river [1]Card) (*Hand, *int, *[5]Card) {
	cards := concatenate(hand, flop, turn, river)
	checkers := []Checker{
		RoyalFlushChecker{},
		StraightFlushChecker{},
		FullHouseChecker{},
		FlushChecker{},
		StraightChecker{},
		ThreeOfAKindChecker{},
		TwoPairChecker{},
		OnePairChecker{},
		HighCardChecker{}}
	for _, c := range checkers {
		h, w, c := c.execute(cards)
		if w > 0 {
			return &h, &w, c
		}
	}
	return nil, nil, nil
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

func fill(start []Card, complete []Card) *[5]Card {
	var result [5]Card
	completeWithoutStart := remove(complete, start)
	for i := 0; i < len(result); i++ {
		if i < len(start) {
			result[i] = start[i]
		} else {
			result[i] = completeWithoutStart[i-len(start)]
		}
	}
	return &result
}

func groupByID(cards [7]Card) map[string][]Card {
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

func groupBySuit(cards [7]Card) map[string][]Card {
	m := make(map[string][]Card)
	for i := 0; i < len(cards); i++ {
		if m[cards[i].Suit] == nil {
			m[cards[i].Suit] = []Card{cards[i]}
		} else {
			m[cards[i].Suit] = append(m[cards[i].Suit], cards[i])
		}
	}
	return m
}

func concatenate(hand [2]Card, flop [3]Card, turn, river [1]Card) [7]Card {
	cards := [7]Card{hand[0], hand[1], flop[0], flop[1], flop[2], turn[0], river[0]}
	sort.Slice(cards[:], func(i, j int) bool {
		return cards[i].Weight > cards[j].Weight
	})
	return cards
}

type Card struct {
	ID     string
	Suit   string
	Weight int
}
