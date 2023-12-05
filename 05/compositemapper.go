package main

type CompositeMapper struct {
	name    string
	mappers []Mapper
}

func (m *CompositeMapper) accepts(num int) bool {
	return true
}

func (m *CompositeMapper) lookup(num int) int {
	for _, mapper := range m.mappers {
		if mapper.accepts(num) {
			return mapper.lookup(num)
		}
	}
	return num
}

func (m *CompositeMapper) add(mapper Mapper) {
	m.mappers = append(m.mappers, mapper)
}
