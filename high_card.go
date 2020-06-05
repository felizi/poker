package main

type HighCardChecker struct {
}

func (o HighCardChecker) execute(avaiableCards AvaiableCards) (Hand, int, *[5]Card) {
	cards := avaiableCards.get()
	var weight int
	var result [5]Card
	for i := 0; i < 5; i++ {
		weight += cards[i].Weight
		result[i] = cards[i]
	}
	return HighCard, weight, &result
}
