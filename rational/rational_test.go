package rational

import "testing"

func TestGenerateRational(t *testing.T) {
	r := NewRational(32, 16)
	if r.Num != 2 {
		t.Errorf("r.Num is expected 2, got=%d", r.Num)
	}
	if r.Dem != 1 {
		t.Errorf("r.Dem is expected 1, got=%d", r.Dem)
	}
	s := NewRational(32, -16)
	if s.Num != -2 {
		t.Errorf("r.Num is expected -2, got=%d", r.Num)
	}
	if s.Dem != 1 {
		t.Errorf("r.Dem is expected 1, got=%d", r.Dem)
	}
}

func TestSumSub(t *testing.T) {
	r := NewRational(3, 4)
	s := NewRational(5, 6)
	result := Sum(r, s)
	if result.Num != 19 {
		t.Errorf("result.Num is expected 19, got=%d", result.Num)
	}
	if result.Dem != 12 {
		t.Errorf("result.Dem is expected 12, got=%d", result.Dem)
	}
	result2 := Sub(r, s)
	if result2.Num != -1 {
		t.Errorf("result.Num is expected -1, got=%d", result2.Num)
	}
	if result2.Dem != 12 {
		t.Errorf("result.Dem is expected 12, got=%d", result2.Dem)
	}
}

func TestProd(t *testing.T) {
	r := NewRational(3, 4)
	s := NewRational(5, 7)
	result := Prod(r, s)
	if result.Num != 15 {
		t.Errorf("result.Num is expected 15, got=%d", result.Num)
	}
	if result.Dem != 28 {
		t.Errorf("result.Dem is expected 28, got=%d", result.Dem)
	}
}

func TestDivide(t *testing.T) {
	r := NewRational(3, 4)
	s := NewRational(5, 7)
	result := Divide(r, s)
	if result.Num != 21 {
		t.Errorf("result.Num is expected 21, got=%d", result.Num)
	}
	if result.Dem != 20 {
		t.Errorf("result.Dem is expected 20, got=%d", result.Dem)
	}
}
