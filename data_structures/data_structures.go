package data_structures

type Container interface {
	Contains(item string) bool
}

type containerImpl struct {
	items []string
}

func (c *containerImpl) Contains(item string) bool {
	for _, i := range c.items {
		if i == item {
			return true
		}
	}

	return false
}

type Stack struct {
	containerImpl
}

func NewStack() *Stack {
	return &Stack{}
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

func (s *Stack) Size() int {
	return len(s.items)
}

type Queue struct {
	containerImpl
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Items() []string {
    return q.items
}

func (q *Queue) Enqueue(item string) {
    q.items = append(q.items, item)
}

func (q *Queue) Dequeue() string {
    item := ""
    if len(q.items) > 0 {
        item = q.items[0]
        q.items = q.items[1:]
    }

    return item
}

func (q *Queue) Size() int {
	return len(q.items)
}
