package main

type FlushChecker struct {
}

func (o FlushChecker) execute(cards [7]Card) (Hand, int, *[5]Card) {
	m := groupBySuit(cards)

	var matches []string
	for k, v := range m {
		if len(v) == 5 {
			matches = append(matches, k)
		}
	}
	if len(matches) == 1 {
		var result [5]Card
		var weight int
		for i := 0; i < len(matches); i++ {
			for x := 0; x < len(m[matches[i]]); x++ {
				c := m[matches[i]][x]
				weight += c.Weight
				result[x] = c
			}
		}
		return Flush, weight, &result
	}

	return Flush, 0, nil
}
