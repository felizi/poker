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
	execute(avaiableCards AvaiableCards) (Hand, int, *[5]Card)
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
	cards = remove(cards, player1[:])

	player2 := hand(cards)
	fmt.Printf("player2: %v\n", player2)
	cards = remove(cards, player2[:])

	flop := flop(cards)
	cards = remove(cards, flop[:])

	turn := turn(cards)
	cards = remove(cards, turn[:])

	river := river(remove(cards, turn[:]))
	cards = remove(cards, river[:])

	fmt.Printf("community cards: %v\n", []Card{flop[0], flop[1], flop[2], turn[0], river[0]})

	ac1 := AvaiableCards{player1, flop, turn, river}
	h1, w1, c1 := check(ac1)
	fmt.Printf("Player1:\t[%v]\t[%v]\t%v\n", *h1, *w1, *c1)

	ac2 := AvaiableCards{player2, flop, turn, river}
	h2, w2, c2 := check(ac2)
	fmt.Printf("Player2:\t[%v]\t[%v]\t%v\n", *h2, *w2, *c2)

	if *h1 == *h2 && *w1 == *w2 {
		fmt.Println("split pot!")
	} else if *h1 > *h2 || (*h1 == *h2 && *w1 > *w2) {
		fmt.Println("Player 1 winner")
	} else if *h1 < *h2 || (*h1 == *h2 && *w1 < *w2) {
		fmt.Println("Player 2 winner")
	} else {
		fmt.Println("fodeu")
	}

}

func check(avaiableCards AvaiableCards) (*Hand, *int, *[5]Card) {
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
		h, w, c := c.execute(avaiableCards)
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
	cards = remove(cards, []Card{x})
	y := random(cards)
	return [2]Card{x, y}
}

func flop(cards []Card) [3]Card {
	x := random(cards)
	cards = remove(cards, []Card{x})
	y := random(cards)
	cards = remove(cards, []Card{y})
	z := random(remove(cards, []Card{y}))
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

type Card struct {
	ID     string
	Suit   string
	Weight int
}

type AvaiableCards struct {
	Hand  [2]Card
	Flop  [3]Card
	Turn  [1]Card
	River [1]Card
}

func (o AvaiableCards) get() [7]Card {
	cards := [7]Card{o.Hand[0], o.Hand[1], o.Flop[0], o.Flop[1], o.Flop[2], o.Turn[0], o.River[0]}
	sort.Slice(cards[:], func(i, j int) bool {
		return cards[i].Weight > cards[j].Weight
	})
	return cards
}

func (o AvaiableCards) groupBySuit() map[string][]Card {
	cards := o.get()
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

func (o AvaiableCards) groupByID() map[string][]Card {
	cards := o.get()
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
