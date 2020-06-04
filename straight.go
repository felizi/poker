package main

type StraightChecker struct {
}

func (o StraightChecker) execute(cards [7]Card) (Hand, int, *[5]Card) {
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
			return Straight, weight, &result
		}
	}

	return Straight, 0, nil
}
