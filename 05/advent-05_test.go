package main

import (
	"testing"
)

func TestSimpleMapper(t *testing.T) {
	mapper := &SimpleMapper{}
	want := 1
	if !mapper.accepts(want) {
		t.Errorf("mapper should accept %d", want)
	}
	got := mapper.lookup(1)
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestRangeMapper(t *testing.T) {
	mapper := &RangeMapper{dst: 1, src: 3, len: 10}
	if mapper.accepts(2) {
		t.Errorf("2 should not be in mapper's range")
	}
	if !mapper.accepts(3) {
		t.Errorf("3 should be in mapper's range")
	}
	if !mapper.accepts(13) {
		t.Errorf("13 should be in mapper's range")
	}
	if mapper.accepts(14) {
		t.Errorf("14 should be in mapper's range")
	}
	want := 5
	got := mapper.lookup(7)
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestCompositeMapper(t *testing.T) {
	mapper := &CompositeMapper{name: "test"}
	mapper.add(&RangeMapper{dst: 50, src: 98, len: 2})
	mapper.add(&RangeMapper{dst: 52, src: 50, len: 48})
	if mapper.name != "test" {
		t.Errorf("mapper has no name")
	}
	if !mapper.accepts(1) {
		t.Errorf("mapper should accept %d", 1)
	}
	if mapper.lookup(79) != 81 {
		t.Errorf("got %d, wanted 81", mapper.lookup(79))
	}
	if mapper.lookup(14) != 14 {
		t.Errorf("got %d, wanted 14", mapper.lookup(79))
	}
	if mapper.lookup(55) != 57 {
		t.Errorf("got %d, wanted 57", mapper.lookup(79))
	}
	if mapper.lookup(13) != 13 {
		t.Errorf("got %d, wanted 13", mapper.lookup(79))
	}
}
