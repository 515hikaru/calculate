package calc

import (
	"testing"

	"github.com/515hikaru/calculate/parser"
)

func TestCalc(t *testing.T) {
	inputTokens := []parser.Token{
		parser.Token{TokenLiteral: "1", TokenType: parser.Numeric},
		parser.Token{TokenLiteral: "2", TokenType: parser.Numeric},
		parser.Token{TokenLiteral: "/", TokenType: parser.Operator},
	}
	c := Calc(inputTokens)
	if c.Num != 1 {
		t.Errorf("c.Num is expected 1, got=%d", c.Num)
	}
	if c.Dem != 2 {
		t.Errorf("c.Dem is expected 2, got=%d", c.Dem)
	}
}

func TestCalcMiddleExpression(t *testing.T) {
	inputTokens := []parser.Token{
		parser.Token{TokenLiteral: "3", TokenType: parser.Numeric},
		parser.Token{TokenLiteral: "4", TokenType: parser.Numeric},
		parser.Token{TokenLiteral: "+", TokenType: parser.Operator},
		parser.Token{TokenLiteral: "1", TokenType: parser.Numeric},
		parser.Token{TokenLiteral: "2", TokenType: parser.Numeric},
		parser.Token{TokenLiteral: "-", TokenType: parser.Operator},
		parser.Token{TokenLiteral: "*", TokenType: parser.Operator},
	}
	c := Calc(inputTokens)
	if c.Num != -7 {
		t.Errorf("c.Num is expected -7, got=%d", c.Num)
	}
	if c.Dem != 1 {
		t.Errorf("c.Dem is expected 1, got=%d", c.Dem)
	}

}
