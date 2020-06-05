package main

import "testing"

func TestOnePairHandAndTurn(t *testing.T) {
	checker := OnePairChecker{}
	avaiableCards := AvaiableCards{
		[2]Card{{"A", "", 7}, {"K", "", 6}},
		[3]Card{{"Q", "", 5}, {"J", "", 4}, {"9", "", 3}},
		[1]Card{{"A", "", 7}},
		[1]Card{{"7", "", 1}},
	}

	hand, _, cards := checker.execute(avaiableCards)

	if hand == OnePair && cards[0].ID == "A" && cards[1].ID == "A" {
		t.Log("Ok")
	} else {
		t.Errorf("Not ok")
	}
}

func TestOnePairInHand(t *testing.T) {
	checker := OnePairChecker{}
	avaiableCards := AvaiableCards{
		[2]Card{{"A", "", 7}, {"A", "", 7}},
		[3]Card{{"Q", "", 5}, {"J", "", 4}, {"9", "", 3}},
		[1]Card{{"K", "", 6}},
		[1]Card{{"7", "", 1}},
	}

	hand, _, cards := checker.execute(avaiableCards)

	if hand == OnePair && cards[0].ID == "A" && cards[1].ID == "A" {
		t.Log("Ok")
	} else {
		t.Errorf("Not ok")
	}
}

func TestOnePairNotExist(t *testing.T) {
	checker := OnePairChecker{}
	avaiableCards := AvaiableCards{
		[2]Card{{"7", "", 1}, {"K", "", 6}},
		[3]Card{{"Q", "", 5}, {"J", "", 4}, {"9", "", 3}},
		[1]Card{{"8", "", 2}},
		[1]Card{{"A", "", 7}},
	}

	hand, _, cards := checker.execute(avaiableCards)

	if hand == OnePair && cards == nil {
		t.Log("Ok")
	} else {
		t.Errorf("Not ok")
	}
}
