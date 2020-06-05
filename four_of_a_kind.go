package main

type FourOfAKindChecker struct {
}

func (o FourOfAKindChecker) execute(avaiableCards AvaiableCards) (Hand, int, *[5]Card) {
	cards := avaiableCards.get()
	m := avaiableCards.groupByID()
	var matches []string
	for k, v := range m {
		if len(v) == 4 {
			matches = append(matches, k)
		}
	}
	if len(matches) == 1 {
		var temp [4]Card
		for i := 0; i < len(matches); i++ {
			for x := 0; x < len(m[matches[i]]); x++ {
				c := m[matches[i]][x]
				temp[x] = c
			}
		}

		result := fill(temp[:], cards[:])
		var weight int
		for i := 0; i < len(result); i++ {
			weight += result[i].Weight
		}

		return FourOfAKind, weight, result
	}

	return FourOfAKind, 0, nil
}
