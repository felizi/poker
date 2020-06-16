package main

type StraightChecker struct {
}

func (o StraightChecker) execute(avaiableCards AvaiableCards) (Hand, int, *[5]Card) {
	cards := avaiableCards.get()
	h, w, c := process(cards)

	if w == 0 {
		var hasAces = false
		for x := 0; x < len(cards); x++ {
			if cards[x].ID == "A" {
				cards[x].Weight = 0
				hasAces = true
			}
		}
		if hasAces {
			cards = avaiableCards.order(cards)

			newH, newW, newC := process(cards)

			if newW > 0 {
				return newH, newW, newC
			}
		}
	}
	return h, w, c
}

func process(cards [7]Card) (Hand, int, *[5]Card) {
	var result [5]Card
	var weight int
	var idx int

	for x := 0; x < len(cards); x++ {
		if idx == 0 || result[idx-1].Weight-1 == cards[x].Weight {
			weight += cards[x].Weight
			result[idx] = cards[x]
			idx++
		} else {
			weight = cards[x].Weight
			result = [5]Card{}
			idx = 0
			result[idx] = cards[x]
			idx++
		}
		if idx >= 5 {
			return Straight, weight, &result
		}
	}
	return Straight, 0, nil
}
