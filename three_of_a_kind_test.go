package main

import "testing"

func TestTreeOfAKindHandFlopRiver(t *testing.T) {
	checker := ThreeOfAKindChecker{}
	avaiableCards := AvaiableCards{
		[2]Card{{"A", "", 12}, {"3", "", 2}},
		[3]Card{{"A", "", 12}, {"K", "", 11}, {"5", "", 4}},
		[1]Card{{"8", "", 7}},
		[1]Card{{"A", "", 12}},
	}

	hand, _, cards := checker.execute(avaiableCards)

	if hand == ThreeOfAKind && cards[0].ID == "A" && cards[1].ID == "A" && cards[2].ID == "A" {
		t.Log("Ok")
	} else {
		t.Errorf("Not ok %v %v", hand, cards)
	}
}

func TestTreeOfAKindFlopTurnRiver(t *testing.T) {
	checker := ThreeOfAKindChecker{}
	avaiableCards := AvaiableCards{
		[2]Card{{"8", "", 7}, {"3", "", 2}},
		[3]Card{{"A", "", 12}, {"K", "", 11}, {"5", "", 4}},
		[1]Card{{"A", "", 12}},
		[1]Card{{"A", "", 12}},
	}

	hand, _, cards := checker.execute(avaiableCards)

	if hand == ThreeOfAKind && cards[0].ID == "A" && cards[1].ID == "A" && cards[2].ID == "A" {
		t.Log("Ok")
	} else {
		t.Errorf("Not ok %v %v", hand, cards)
	}
}

func TestTreeOfAKindFlop(t *testing.T) {
	checker := ThreeOfAKindChecker{}
	avaiableCards := AvaiableCards{
		[2]Card{{"8", "", 7}, {"3", "", 2}},
		[3]Card{{"A", "", 12}, {"A", "", 12}, {"A", "", 12}},
		[1]Card{{"K", "", 11}},
		[1]Card{{"5", "", 4}},
	}

	hand, _, cards := checker.execute(avaiableCards)

	if hand == ThreeOfAKind && cards[0].ID == "A" && cards[1].ID == "A" && cards[2].ID == "A" {
		t.Log("Ok")
	} else {
		t.Errorf("Not ok %v %v", hand, cards)
	}
}

func TestTreeOfAKindNotExists(t *testing.T) {
	checker := ThreeOfAKindChecker{}
	avaiableCards := AvaiableCards{
		[2]Card{{"8", "", 7}, {"3", "", 2}},
		[3]Card{{"6", "", 5}, {"K", "", 11}, {"5", "", 4}},
		[1]Card{{"A", "", 12}},
		[1]Card{{"A", "", 12}},
	}

	hand, _, cards := checker.execute(avaiableCards)

	if hand == ThreeOfAKind && cards == nil {
		t.Log("Ok")
	} else {
		t.Errorf("Not ok %v %v", hand, cards)
	}
}
