package parser

import (
	"strings"
)

// token types
const (
	Operator = "operator"
	Numeric  = "numeric"
)

// operators
const (
	Plus   = "+"
	Minus  = "-"
	Prod   = "*"
	Divide = "/"
)

// ignore token
const (
	space = " "
	empty = ""
)

type Token struct {
	TokenLiteral string
	TokenType    string
}

func Parse(s string) []Token {
	split := strings.Split(s, space)
	tks := make([]Token, 0)

	for _, c := range split {
		if c == empty {
			continue
		}

		tk := Token{}
		tk.TokenLiteral = c

		switch c {
		case Plus, Minus, Prod, Divide:
			tk.TokenType = Operator
		default:
			tk.TokenType = Numeric
		}

		tks = append(tks, tk)
	}

	return tks
}
