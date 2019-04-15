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
	absNum := int(math.Abs(float64(num)))
	g := gcd(absNum, dem)
	return Rational{
		Num: num / g,
		Dem: dem / g,
	}
}

func Prod(r, s Rational) Rational {
	newNum := r.Num * s.Num
	newDem := r.Dem * s.Dem
	return NewRational(newNum, newDem)
}
