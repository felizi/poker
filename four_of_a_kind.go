package main

type FourOfAKindChecker struct {
}

func (o FourOfAKindChecker) execute(cards [7]Card) (Hand, int, *[5]Card) {
	m := groupByID(cards)
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
		return FourOfAKind, weight, fill(result[:], cards[:])
	}

	return FourOfAKind, 0, nil
}
