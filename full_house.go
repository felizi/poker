package main

type FullHouseChecker struct {
}

func (o FullHouseChecker) execute(cards [7]Card) (Hand, int, *[5]Card) {
	m := groupByID(cards)

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
		cards := append(threePair, twoPair...)
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
		return FullHouse, weight, &result
	}

	return FullHouse, 0, nil
}
