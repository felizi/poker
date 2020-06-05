package main

import (
	"sort"
)

type TwoPairChecker struct {
}

func (o TwoPairChecker) execute(avaiableCards AvaiableCards) (Hand, int, *[5]Card) {
	cards := avaiableCards.get()
	m := avaiableCards.groupByID()
	var matches []string
	for k, v := range m {
		if len(v) == 2 {
			matches = append(matches, k)

		}
	}
	if len(matches) == 2 {
		var temp [4]Card
		var idx int
		for i := 0; i < len(matches); i++ {
			for x := 0; x < len(m[matches[i]]); x++ {
				c := m[matches[i]][x]
				temp[idx] = c
				idx++
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
		return TwoPair, weight, &result
	}

	return TwoPair, 0, nil
}
