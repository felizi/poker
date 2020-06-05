package main

import "sort"

type OnePairChecker struct {
}

func (o OnePairChecker) execute(avaiableCards AvaiableCards) (Hand, int, *[5]Card) {
	cards := avaiableCards.get()
	m := avaiableCards.groupByID()
	var matches []string
	for k, v := range m {
		if len(v) == 2 {
			matches = append(matches, k)
		}
	}
	if len(matches) == 1 {
		var temp [2]Card

		for i := 0; i < len(matches); i++ {
			for x := 0; x < len(m[matches[i]]); x++ {
				c := m[matches[i]][x]
				temp[x] = c
			}
		}

		sort.Slice(temp[:], func(i, j int) bool {
			return temp[i].Weight > temp[j].Weight
		})

		result := fill(temp[:], cards[:])
		var weight int
		for i := 0; i < len(result); i++ {
			weight += result[i].Weight
		}
		return OnePair, weight, &result
	}

	return OnePair, 0, nil
}
