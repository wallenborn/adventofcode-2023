package main

import (
	"testing"
)

func TestCanCountOneHash(t *testing.T) {
	want := 1
	got := Count("#", []int{1})
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestCanCountOneDot(t *testing.T) {
	want := 1
	got := Count(".", []int{})
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestCanCountTrailingDots(t *testing.T) {
	want := 1
	got := Count("#..", []int{1})
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestCanCountTrailingQuestionMarks(t *testing.T) {
	want := 1
	got := Count("#??", []int{1})
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestCanCountAllQuestionMarks(t *testing.T) {
	want := 3
	got := Count("?????", []int{2, 1})
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestCanCountOneQuestionmarkAsHash(t *testing.T) {
	want := 1
	got := Count("?", []int{1})
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestCanCountOneQuestionmarkAsDot(t *testing.T) {
	want := 1
	got := Count("?", []int{})
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestCanCountTwoHashes(t *testing.T) {
	want := 1
	got := Count("##", []int{2})
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestString0(t *testing.T) {
	want := 1
	got := Count("#.#", []int{1, 1})
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestLine1(t *testing.T) {
	want := 1
	got := Count("#.#.###", []int{1, 1, 3})
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestLine2(t *testing.T) {
	want := 4
	got := Count(".??..??...?##.", []int{1, 1, 3})
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestLine3(t *testing.T) {
	want := 1
	got := Count("?#?#?#?#?#?#?#?", []int{1, 3, 1, 6})
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestLine4(t *testing.T) {
	want := 1
	got := Count("????.#...#...", []int{4, 1, 1})
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestLine5(t *testing.T) {
	want := 4
	got := Count("????.######..#####.", []int{1, 6, 5})
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestLine6(t *testing.T) {
	want := 10
	got := Count("?###????????", []int{3, 2, 1})
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
