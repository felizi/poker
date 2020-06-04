package main

type TwoPairChecker struct {
}

func (o TwoPairChecker) execute(cards [7]Card) (Hand, int, *[5]Card) {
	m := groupByID(cards)
	var matches []string
	for k, v := range m {
		if len(v) == 2 {
			matches = append(matches, k)
		}
	}
	if len(matches) == 2 {
		var result [4]Card
		var weight int
		var idx int
		for i := 0; i < len(matches); i++ {
			for x := 0; x < len(m[matches[i]]); x++ {
				c := m[matches[i]][x]
				weight += c.Weight
				result[idx] = c
				idx++
			}
		}
		return TwoPair, weight, fill(result[:], cards[:])
	}

	return TwoPair, 0, nil
}
