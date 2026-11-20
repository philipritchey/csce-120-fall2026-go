package main

type IntStack struct {
	items []int
}

func (s *IntStack) Push(item int) {
	s.items = append(s.items, item)
}

func (s *IntStack) Pop() (int, bool) {
	if len(s.items) == 0 {
		return 0, false
	}

	lastIndex := len(s.items) - 1
	item := s.items[lastIndex]
	s.items = s.items[:lastIndex]
	return item, true
}

func (s *IntStack) IsEmpty() bool {
	return len(s.items) == 0
}
