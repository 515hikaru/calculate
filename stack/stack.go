package stack

type Stack struct {
	idx   int
	items []string
}

func NewStack() Stack {
	return Stack{}
}

func (s *Stack) Push(x string) {
	s.items = append(s.items, x)
	s.idx += 1
}

func (s *Stack) Pop() string {
	s.idx -= 1
	return s.items[s.idx]
}
