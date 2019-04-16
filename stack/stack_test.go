package stack

import "testing"

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
	s.Push("foo")
	s.Push("bar")
	if s.idx != 2 {
		t.Errorf("s.idx is expected 2. got=%d", s.idx)
	}
	if s.items[0] != "foo" {
		t.Errorf("s.items[0] is expected foo, got=%v", s.items[0])
	}
	if s.items[1] != "bar" {
		t.Errorf("s.items[1] is expected bar, got=%v", s.items[1])
	}
}

func TestStackPop(t *testing.T) {
	s := NewStack()
	s.Push("foo")
	s.Push("bar")
	v := s.Pop()
	if v != "bar" {
		t.Errorf("v is expected bar, got=%v", v)
	}
	if s.idx != 1 {
		t.Errorf("s.idx is expected 1. got=%d", s.idx)
	}
}
