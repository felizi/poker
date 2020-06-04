package main

type ThreeOfAKindChecker struct {
}

func (o ThreeOfAKindChecker) execute(cards [7]Card) (Hand, int, *[5]Card) {
	m := groupByID(cards)
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
		return ThreeOfAKind, weight, fill(result[:], cards[:])
	}

	return ThreeOfAKind, 0, nil
}
