package main

import "fmt"

type RangeMapper struct {
	dst int
	src int
	len int
}

func (m *RangeMapper) accepts(num int) bool {
	return num >= m.src && num < m.src+m.len
}

func (m *RangeMapper) lookup(num int) int {
	if m.accepts(num) {
		return num + m.dst - m.src
	} else {
		msg := fmt.Sprintf("Number %d out of range [%d,%d]", num, m.src, m.src+m.len)
		panic(msg)
	}
}
