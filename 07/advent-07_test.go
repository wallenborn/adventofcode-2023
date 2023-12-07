package main

import (
	"testing"
)

func TestHighCard(t *testing.T) {
	if !isHighCard("23456") {
		t.Errorf("23456 should be HighCard")
	}
	if isHighCard("22456") {
		t.Errorf("22456 should be OnePair")
	}
}

func TestOnePair(t *testing.T) {
	if isOnePair("23456") {
		t.Errorf("23456 should be HighCard")
	}
	if !isOnePair("22456") {
		t.Errorf("22456 should be OnePair")
	}
}

func TestTwoPair(t *testing.T) {
	if isTwoPair("23456") {
		t.Errorf("23456 should be HighCard")
	}
	if !isTwoPair("22446") {
		t.Errorf("22446 should be TwoPair")
	}
}

func TestThreeOfAKind(t *testing.T) {
	if isThreeOfAKind("23456") {
		t.Errorf("23456 should be HighCard")
	}
	if !isThreeOfAKind("22246") {
		t.Errorf("22246 should be ThreeOfAKind")
	}
}

func TestFullHouse(t *testing.T) {
	if isFullHouse("23456") {
		t.Errorf("23456 should be HighCard")
	}
	if !isFullHouse("22244") {
		t.Errorf("22244 should be FullHouse")
	}
}

func TestFourOfAKind(t *testing.T) {
	if isFourOfAKind("23456") {
		t.Errorf("23456 should be HighCard")
	}
	if !isFourOfAKind("22226") {
		t.Errorf("22226 should be FourOfAKind")
	}
}

func TestFiveOfAKind(t *testing.T) {
	if isFiveOfAKind("23456") {
		t.Errorf("23456 should be HighCard")
	}
	if !isFiveOfAKind("22222") {
		t.Errorf("22222 should be FiveOfAKind")
	}
}
