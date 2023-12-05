package main

type SimpleMapper struct {
}

func (s *SimpleMapper) accepts(num int) bool {
	return true
}

func (s *SimpleMapper) lookup(num int) int {
	return num
}
