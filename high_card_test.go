package main

import "testing"

func TestHighCardAces(t *testing.T) {
	checker := HighCardChecker{}
	avaiable := [7]Card{
		{
			ID:     "A",
			Suit:   "",
			Weight: 0,
		},
	}
	hand, _, cards := checker.execute(avaiable)

	if hand == HighCard && cards[0].ID == "A" {
		t.Log("Ok")
	} else {
		t.Errorf("Not ok")
	}

}
