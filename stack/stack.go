package stack

import "github.com/515hikaru/calculate/rational"

type Stack struct {
	idx   int
	items []rational.Rational
}

func NewStack() Stack {
	return Stack{}
}

func (s *Stack) Push(x rational.Rational) {
	if len(s.items) == s.idx {
		s.items = append(s.items, x)
	} else {
		s.items[s.idx] = x
	}
	s.idx += 1
}

func (s *Stack) Pop() rational.Rational {
	s.idx -= 1
	return s.items[s.idx]
}
