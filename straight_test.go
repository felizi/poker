package main

import (
	"testing"
)

func TestStraightHandFlopRiver(t *testing.T) {
	checker := StraightChecker{}
	avaiableCards := AvaiableCards{
		[2]Card{{"A", "", 12}, {"10", "", 8}},
		[3]Card{{"Q", "", 10}, {"K", "", 11}, {"10", "", 8}},
		[1]Card{{"J", "", 9}},
		[1]Card{{"9", "", 7}},
	}

	hand, _, cards := checker.execute(avaiableCards)
	if hand == Straight && cards[0].ID == "A" && cards[4].ID == "10" {
		t.Log("Ok")
	} else {
		t.Errorf("Not ok %v %v", hand, cards)
	}
}

func TestStraightFlopTurnRiver(t *testing.T) {
	checker := StraightChecker{}
	avaiableCards := AvaiableCards{
		[2]Card{{"9", "", 7}, {"5", "", 3}},
		[3]Card{{"Q", "", 10}, {"K", "", 11}, {"10", "", 8}},
		[1]Card{{"J", "", 9}},
		[1]Card{{"A", "", 12}},
	}

	hand, _, cards := checker.execute(avaiableCards)
	if hand == Straight && cards[0].ID == "A" && cards[4].ID == "10" {
		t.Log("Ok")
	} else {
		t.Errorf("Not ok %v %v", hand, cards)
	}
}

func TestStraightAllCards(t *testing.T) {
	checker := StraightChecker{}
	avaiableCards := AvaiableCards{
		[2]Card{{"2", "", 0}, {"5", "", 3}},
		[3]Card{{"3", "", 1}, {"7", "", 5}, {"8", "", 6}},
		[1]Card{{"4", "", 2}},
		[1]Card{{"6", "", 4}},
	}

	hand, _, cards := checker.execute(avaiableCards)
	if hand == Straight && cards[0].ID == "8" && cards[4].ID == "4" {
		t.Log("Ok")
	} else {
		t.Errorf("Not ok %v %v", hand, cards)
	}
}
func TestStraightLowerWithAces(t *testing.T) {
	checker := StraightChecker{}
	avaiableCards := AvaiableCards{
		[2]Card{{"2", "", 1}, {"5", "", 4}},
		[3]Card{{"3", "", 2}, {"7", "", 6}, {"8", "", 7}},
		[1]Card{{"4", "", 3}},
		[1]Card{{"A", "", 13}},
	}

	hand, _, cards := checker.execute(avaiableCards)

	if hand == Straight && cards[0].ID == "5" && cards[4].ID == "A" {
		t.Log("Ok")
	} else {
		t.Errorf("Not ok %v %v", hand, cards)
	}
}

func TestStraightOutOfOrderLow(t *testing.T) {
	checker := StraightChecker{}
	avaiableCards := AvaiableCards{
		[2]Card{{"5", "", 4}, {"9", "", 8}},
		[3]Card{{"7", "", 6}, {"8", "", 7}, {"J", "", 10}},
		[1]Card{{"K", "", 12}},
		[1]Card{{"6", "", 5}},
	}

	hand, _, cards := checker.execute(avaiableCards)

	if hand == Straight && cards[0].ID == "9" && cards[4].ID == "5" {
		t.Log("Ok")
	} else {
		t.Errorf("Not ok %v %v", hand, cards)
	}
}

func TestStraightOutOfOrderHigh(t *testing.T) {
	checker := StraightChecker{}
	avaiableCards := AvaiableCards{
		[2]Card{{"5", "", 4}, {"9", "", 8}},
		[3]Card{{"7", "", 6}, {"8", "", 7}, {"J", "", 10}},
		[1]Card{{"10", "", 9}},
		[1]Card{{"6", "", 5}},
	}

	hand, _, cards := checker.execute(avaiableCards)

	if hand == Straight && cards[0].ID == "J" && cards[4].ID == "7" {
		t.Log("Ok")
	} else {
		t.Errorf("Not ok %v %v", hand, cards)
	}
}
