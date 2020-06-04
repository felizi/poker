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

	check("Player1", player1, flop, turn, river)
	check("Player2", player2, flop, turn, river)
}

func check(name string, hand [2]Card, flop [3]Card, turn, river [1]Card) {
	fmt.Printf("%v: \t %v\n", name, concatenate(hand, flop, turn, river))
	w, rf := royalFlush(hand, flop, turn, river)
	fmt.Printf("%v - royal flush: \t %v %v\n", name, w, rf)
	w, sf := straightFlush(hand, flop, turn, river)
	fmt.Printf("%v - straight flush: \t %v %v\n", name, w, sf)
	w, foak := fourOfAKind(hand, flop, turn, river)
	fmt.Printf("%v - four of a kind: \t %v %v\n", name, w, foak)
	w, fh := fullHouse(hand, flop, turn, river)
	fmt.Printf("%v - full house: \t %v %v\n", name, w, fh)
	w, f := flush(hand, flop, turn, river)
	fmt.Printf("%v - flush: \t %v %v\n", name, w, f)
	w, s := straight(hand, flop, turn, river)
	fmt.Printf("%v - straight: \t %v %v\n", name, w, s)
	w, t := threeOfAKind(hand, flop, turn, river)
	fmt.Printf("%v - three of a kind: \t %v %v\n", name, w, t)
	w, p := twoPair(hand, flop, turn, river)
	fmt.Printf("%v - two pair: \t %v %v\n", name, w, p)
	w, c := onePair(hand, flop, turn, river)
	fmt.Printf("%v - one pair: \t %v %v\n", name, w, c)
	w, h := highCard(hand, flop, turn, river)
	fmt.Printf("%v - high card: \t %v  %v\n", name, w, h)
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

func royalFlush(hand [2]Card, flop [3]Card, turn, river [1]Card) (int, *[5]Card) {
	cards := concatenate(hand, flop, turn, river)

	var result [5]Card
	var weight int
	var lastWeight int
	var idx int
	var suit string
	for x := 0; x < len(cards); x++ {
		if idx == 0 && cards[x].ID == "A" || lastWeight-1 == cards[x].Weight && suit == cards[x].Suit {
			lastWeight = cards[x].Weight
			weight += cards[x].Weight
			result[idx] = cards[x]
			idx++
			suit = cards[x].Suit
		} else {
			lastWeight = 0
			weight = 0
			result = [5]Card{}
			idx = 0
			suit = ""
		}
		if idx >= 5 {
			return weight, &result
		}
	}

	return 0, nil
}

func straightFlush(hand [2]Card, flop [3]Card, turn, river [1]Card) (int, *[5]Card) {
	cards := concatenate(hand, flop, turn, river)

	var result [5]Card
	var weight int
	var lastWeight int
	var idx int
	var suit string
	for x := 0; x < len(cards); x++ {
		if x == 0 || lastWeight-1 == cards[x].Weight && suit == cards[x].Suit {
			lastWeight = cards[x].Weight
			weight += cards[x].Weight
			result[idx] = cards[x]
			idx++
			suit = cards[x].Suit
		} else {
			lastWeight = 0
			weight = 0
			result = [5]Card{}
			idx = 0
			suit = ""
		}
		if idx >= 5 {
			return weight, &result
		}
	}

	return 0, nil
}

func fourOfAKind(hand [2]Card, flop [3]Card, turn, river [1]Card) (int, *[4]Card) {
	m := groupByID(hand, flop, turn, river)
	var matches []string
	for k, v := range m {
		if len(v) == 4 {
			matches = append(matches, k)
		}
	}
	if len(matches) == 1 {
		var result [4]Card
		var weight int
		for i := 0; i < len(matches); i++ {
			for x := 0; x < len(m[matches[i]]); x++ {
				c := m[matches[i]][x]
				weight += c.Weight
				result[x] = c
			}
		}
		return weight, &result
	}

	return 0, nil
}

func fullHouse(hand [2]Card, flop [3]Card, turn, river [1]Card) (int, *[5]Card) {
	m := groupByID(hand, flop, turn, river)
	var threePair []string
	var twoPair []string
	for k, v := range m {
		if len(v) == 3 {
			threePair = append(threePair, k)
		}
		if len(v) == 2 {
			twoPair = append(twoPair, k)
		}
	}
	if len(twoPair) == 1 && len(threePair) == 1 {
		cards := append(twoPair, threePair...)
		var result [5]Card
		var weight int
		var idx int
		for i := 0; i < len(cards); i++ {
			for x := 0; x < len(m[cards[i]]); x++ {
				c := m[cards[i]][x]
				weight += c.Weight
				result[idx] = c
				idx++
			}
		}
		return weight, &result
	}

	return 0, nil
}

func flush(hand [2]Card, flop [3]Card, turn, river [1]Card) (int, *[5]Card) {
	m := groupBySuit(hand, flop, turn, river)

	var matches []string
	for k, v := range m {
		if len(v) == 5 {
			matches = append(matches, k)
		}
	}
	if len(matches) == 1 {
		var result [5]Card
		var weight int
		for i := 0; i < len(matches); i++ {
			for x := 0; x < len(m[matches[i]]); x++ {
				c := m[matches[i]][x]
				weight += c.Weight
				result[x] = c
			}
		}
		return weight, &result
	}

	return 0, nil

}

func straight(hand [2]Card, flop [3]Card, turn, river [1]Card) (int, *[5]Card) {
	cards := concatenate(hand, flop, turn, river)

	var result [5]Card
	var weight int
	var lastWeight int
	var idx int
	for x := 0; x < len(cards); x++ {
		if x == 0 || lastWeight-1 == cards[x].Weight {
			lastWeight = cards[x].Weight
			weight += cards[x].Weight
			result[idx] = cards[x]
			idx++
		} else {
			lastWeight = 0
			weight = 0
			result = [5]Card{}
			idx = 0
		}
		if idx >= 5 {
			return weight, &result
		}
	}

	return 0, nil
}

func threeOfAKind(hand [2]Card, flop [3]Card, turn, river [1]Card) (int, *[3]Card) {
	m := groupByID(hand, flop, turn, river)
	var matches []string
	for k, v := range m {
		if len(v) == 3 {
			matches = append(matches, k)
		}
	}
	if len(matches) == 1 {
		var result [3]Card
		var weight int
		for i := 0; i < len(matches); i++ {
			for x := 0; x < len(m[matches[i]]); x++ {
				c := m[matches[i]][x]
				weight += c.Weight
				result[x] = c
			}
		}
		return weight, &result
	}

	return 0, nil
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
	m := groupByID(hand, flop, turn, river)
	var matches []string
	for k, v := range m {
		if len(v) == 2 {
			matches = append(matches, k)
		}
	}
	if len(matches) == 1 {
		var result [2]Card
		var weight int
		for i := 0; i < len(matches); i++ {
			for x := 0; x < len(m[matches[i]]); x++ {
				c := m[matches[i]][x]
				weight += c.Weight
				result[x] = c
			}
		}
		return weight, &result
	}

	return 0, nil
}

func highCard(hand [2]Card, flop [3]Card, turn, river [1]Card) (int, *[5]Card) {
	cards := concatenate(hand, flop, turn, river)
	var sum int
	var result [5]Card
	for i := 0; i < 5; i++ {
		sum += cards[i].Weight
		result[i] = cards[i]
	}
	return sum, &result
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

func groupBySuit(hand [2]Card, flop [3]Card, turn, river [1]Card) map[string][]Card {
	cards := concatenate(hand, flop, turn, river)
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
