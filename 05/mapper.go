package main

type Mapper interface {
	accepts(int) bool
	lookup(int) int
}
