package rational

import "math"

type Rational struct {
	Num int
	Dem int
}

func gcd(a, b int) int {
	if a < b {
		a, b = b, a
	}
	if a%b == 0 {
		return b
	}
	return gcd(b, a%b)
}

func NewRational(num, dem int) Rational {
	// denominator should be non-negative
	if dem < 0 {
		dem = -dem
		num = -num
	}

	r := Rational{
		Num: num,
		Dem: dem,
	}
	r.normalize()
	return r
}

func (r *Rational) normalize() {
	absNum := int(math.Abs(float64(r.Num)))
	g := gcd(absNum, r.Dem)
	r.Num = r.Num / g
	r.Dem = r.Dem / g
}

func (r Rational) Inverse() Rational {
	return NewRational(r.Dem, r.Num)
}

func Prod(r, s Rational) Rational {
	newNum := r.Num * s.Num
	newDem := r.Dem * s.Dem
	return NewRational(newNum, newDem)
}

func Divide(r, s Rational) Rational {
	return Prod(r, s.Inverse())
}
