package main

import "testing"

func TestTwoPairHandAndFlop(t *testing.T) {
	checker := TwoPairChecker{}
	avaiableCards := AvaiableCards{
		[2]Card{{"A", "", 12}, {"K", "", 11}},
		[3]Card{{"Q", "", 5}, {"K", "", 11}, {"A", "", 12}},
		[1]Card{{"8", "", 7}},
		[1]Card{{"7", "", 1}},
	}

	hand, _, cards := checker.execute(avaiableCards)

	if hand == TwoPair && cards[0].ID == "A" && cards[2].ID == "K" {
		t.Log("Ok")
	} else {
		t.Errorf("Not ok %v %v", hand, cards)
	}
}

func TestPairFlopAndRiver(t *testing.T) {
	checker := TwoPairChecker{}
	avaiableCards := AvaiableCards{
		[2]Card{{"A", "", 12}, {"Q", "", 10}},
		[3]Card{{"A", "", 12}, {"J", "", 9}, {"9", "", 8}},
		[1]Card{{"K", "", 11}},
		[1]Card{{"Q", "", 10}},
	}

	hand, _, cards := checker.execute(avaiableCards)

	if hand == TwoPair && cards[0].ID == "A" && cards[2].ID == "Q" {
		t.Log("Ok")
	} else {
		t.Errorf("Not ok %v %v", hand, cards)
	}
}

func TestPairFailed(t *testing.T) {
	checker := TwoPairChecker{}
	avaiableCards := AvaiableCards{
		[2]Card{{"5", "", 4}, {"Q", "", 10}},
		[3]Card{{"A", "", 12}, {"J", "", 9}, {"9", "", 8}},
		[1]Card{{"K", "", 11}},
		[1]Card{{"Q", "", 10}},
	}

	hand, _, cards := checker.execute(avaiableCards)

	if hand == TwoPair && cards == nil {
		t.Log("Ok")
	} else {
		t.Errorf("Not ok %v %v", hand, cards)
	}
}
