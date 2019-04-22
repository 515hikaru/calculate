package parser

import (
	"strings"
)

// token types
const (
	operator = "operator"
	numeric  = "numeric"
)

// operators
const (
	plus   = "+"
	minus  = "-"
	prod   = "*"
	divide = "/"
)

// ignore token
const (
	space  = " "
	empty  = ""
)

type token struct {
	tokenLiteral string
	tokenType    string
}

func Parse(s string) []token {
	split := strings.Split(s, space)
	tks := make([]token, 0)

	for _, c := range split {
		if c == empty {
			continue
		}

		tk := token{}
		tk.tokenLiteral = c

		switch c {
		case plus, minus, prod, divide:
			tk.tokenType = operator
		default:
			tk.tokenType = numeric
		}

		tks = append(tks, tk)
	}

	return tks
}
