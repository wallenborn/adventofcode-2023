package main

import (
	"testing"
)

func TestCanParse0(t *testing.T) {
	want := 0
	got, err := StringToInteger("0")
	if got != want || err != nil {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestCanParseZero(t *testing.T) {
	want := 0
	got, err := StringToInteger("zero")
	if got != want || err != nil {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestCanParse1(t *testing.T) {
	want := 1
	got, err := StringToInteger("1")
	if got != want || err != nil {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestCanParseOne(t *testing.T) {
	want := 1
	got, err := StringToInteger("one")
	if got != want || err != nil {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestCanParse2(t *testing.T) {
	want := 2
	got, err := StringToInteger("2")
	if got != want || err != nil {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestCanParseTwo(t *testing.T) {
	want := 2
	got, err := StringToInteger("two")
	if got != want || err != nil {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestCanParse3(t *testing.T) {
	want := 3
	got, err := StringToInteger("3")
	if got != want || err != nil {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestCanParseThree(t *testing.T) {
	want := 3
	got, err := StringToInteger("three")
	if got != want || err != nil {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestCanParse4(t *testing.T) {
	want := 4
	got, err := StringToInteger("4")
	if got != want || err != nil {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestCanParseFour(t *testing.T) {
	want := 4
	got, err := StringToInteger("four")
	if got != want || err != nil {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestCanParse5(t *testing.T) {
	want := 5
	got, err := StringToInteger("5")
	if got != want || err != nil {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestCanParseFive(t *testing.T) {
	want := 5
	got, err := StringToInteger("five")
	if got != want || err != nil {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestCanParse6(t *testing.T) {
	want := 6
	got, err := StringToInteger("6")
	if got != want || err != nil {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestCanParseSix(t *testing.T) {
	want := 6
	got, err := StringToInteger("six")
	if got != want || err != nil {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestCanParse7(t *testing.T) {
	want := 7
	got, err := StringToInteger("7")
	if got != want || err != nil {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestCanParseSeven(t *testing.T) {
	want := 7
	got, err := StringToInteger("seven")
	if got != want || err != nil {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestCanParse8(t *testing.T) {
	want := 8
	got, err := StringToInteger("8")
	if got != want || err != nil {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestCanParseEight(t *testing.T) {
	want := 8
	got, err := StringToInteger("eight")
	if got != want || err != nil {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestCanParse9(t *testing.T) {
	want := 9
	got, err := StringToInteger("9")
	if got != want || err != nil {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestCanParseNine(t *testing.T) {
	want := 9
	got, err := StringToInteger("nine")
	if got != want || err != nil {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
