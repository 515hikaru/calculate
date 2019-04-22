package stack

import (
	"testing"

	"github.com/515hikaru/calculate/rational"
)

func TestStack(t *testing.T) {
	s := NewStack()
	if s.idx != 0 {
		t.Errorf("s.idx is expected 0. got=%d", s.idx)
	}
	if len(s.items) != 0 {
		t.Errorf("s.items is expected empty. got=%v", s.items)
	}
}

func TestStackPush(t *testing.T) {
	s := NewStack()
	s.Push(rational.NewRational(2, 1))
	s.Push(rational.NewRational(1, 1))
	if s.idx != 2 {
		t.Errorf("s.idx is expected 2. got=%d", s.idx)
	}
	if s.items[0].Num != 2 && s.items[0].Dem != 1 {
		t.Errorf("s.items[0] is expected 2, got=%v", s.items[0])
	}
	if s.items[1].Num != 1 && s.items[1].Dem != 1 {
		t.Errorf("s.items[1] is expected bar, got=%v", s.items[1])
	}
}

func TestStackPop(t *testing.T) {
	s := NewStack()
	s.Push(rational.NewRational(2, 1))
	s.Push(rational.NewRational(1, 1))
	v := s.Pop()
	if v.Num != 1 && v.Dem != 1 {
		t.Errorf("v is expected bar, got=%v", v)
	}
	if s.idx != 1 {
		t.Errorf("s.idx is expected 1. got=%d", s.idx)
	}
}
