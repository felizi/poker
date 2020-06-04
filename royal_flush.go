package main

type RoyalFlushChecker struct {
}

func (o RoyalFlushChecker) execute(cards [7]Card) (Hand, int, *[5]Card) {
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
			return RoyalFlush, weight, &result
		}
	}

	return RoyalFlush, 0, nil
}
