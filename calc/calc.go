package calc

import (
	"errors"
	"strconv"

	"github.com/515hikaru/calculate/parser"
	"github.com/515hikaru/calculate/rational"
	"github.com/515hikaru/calculate/stack"
)

func initializeStack() stack.Stack {
	return stack.NewStack()
}

func calc(operator string, r, s rational.Rational) (rational.Rational, error) {
	switch operator {
	case parser.Plus:
		return rational.Sum(r, s), nil
	case parser.Minus:
		return rational.Sub(r, s), nil
	case parser.Prod:
		return rational.Prod(r, s), nil
	case parser.Divide:
		return rational.Divide(r, s), nil
	}
	return rational.Rational{}, errors.New("Unexpected expression")
}

func Calc(tokens []parser.Token) rational.Rational {
	st := stack.NewStack()
	for _, tk := range tokens {
		if tk.TokenType == parser.Operator {
			r := st.Pop()
			s := st.Pop()
			t, err := calc(tk.TokenLiteral, s, r)
			if err != nil {
				panic(err)
			}
			st.Push(t)
		} else {
			v, err := strconv.Atoi(tk.TokenLiteral)
			if err != nil {
				panic(err)
			}
			tmp := rational.NewRational(v, 1)
			st.Push(tmp)
		}

	}
	return st.Pop()
}
