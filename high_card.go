package main

type HighCardChecker struct {
}

func (o HighCardChecker) execute(cards [7]Card) (Hand, int, *[5]Card) {
	var weight int
	var result [5]Card
	for i := 0; i < 5; i++ {
		weight += cards[i].Weight
		result[i] = cards[i]
	}
	return HighCard, weight, &result
}
