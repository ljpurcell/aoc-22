package data_structures

type Stack struct {
	items []string
}

func (s *Stack) Push(item string) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() string {
	item := ""
	if len(s.items) > 0 {
		last := len(s.items) - 1
		item = s.items[last]
		s.items = s.items[:last]
	}

	return item
}

func NewStack() *Stack {
    return &Stack{}
}

func (s *Stack) Size() int {
    return len(s.items)
}
