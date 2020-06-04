package main

type OnePairChecker struct {
}

func (o OnePairChecker) execute(cards [7]Card) (Hand, int, *[5]Card) {
	m := groupByID(cards)
	var matches []string
	for k, v := range m {
		if len(v) == 2 {
			matches = append(matches, k)
			break
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
		return OnePair, weight, fill(result[:], cards[:])
	}

	return OnePair, 0, nil
}
