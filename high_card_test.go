package main

import "testing"

func TestHighCardAcesOnHand(t *testing.T) {
	checker := HighCardChecker{}
	avaiableCards := AvaiableCards{
		[2]Card{{"A", "", 7}, {"K", "", 6}},
		[3]Card{{"Q", "", 5}, {"J", "", 4}, {"9", "", 3}},
		[1]Card{{"8", "", 2}},
		[1]Card{{"7", "", 1}},
	}

	hand, _, cards := checker.execute(avaiableCards)

	if hand == HighCard && cards[0].ID == "A" {
		t.Log("Ok")
	} else {
		t.Errorf("Not ok")
	}
}

func TestHighCardAcesOnRiver(t *testing.T) {
	checker := HighCardChecker{}
	avaiableCards := AvaiableCards{
		[2]Card{{"7", "", 1}, {"K", "", 6}},
		[3]Card{{"Q", "", 5}, {"J", "", 4}, {"9", "", 3}},
		[1]Card{{"8", "", 2}},
		[1]Card{{"A", "", 7}},
	}

	hand, _, cards := checker.execute(avaiableCards)

	if hand == HighCard && cards[0].ID == "A" {
		t.Log("Ok")
	} else {
		t.Errorf("Not ok")
	}
}
