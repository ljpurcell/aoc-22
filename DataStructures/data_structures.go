package datastructures

type Stack struct {
	items []string
}

func (s *Stack) push(item string) {
	s.items = append(s.items, item)
}

func (s *Stack) pop() string {
	item := ""
	if len(s.items) > 0 {
		last := len(s.items) - 1
		item = s.items[last]
		s.items = s.items[:last]
	}

	return item
}
