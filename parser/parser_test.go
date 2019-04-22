package parser

import "testing"

func TestParseSimpleToken(t *testing.T) {
	stringExample := "1 1 +"
	results := Parse(stringExample)
	expected := []Token{
		Token{TokenLiteral: "1", TokenType: Numeric},
		Token{TokenLiteral: "1", TokenType: Numeric},
		Token{TokenLiteral: "+", TokenType: Operator},
	}
	testParser(t, results, expected)

}

func TestParseNegativeaInteger(t *testing.T) {
	inputString := "10 -20 +"
	results := Parse(inputString)
	expected := []Token{
		Token{TokenLiteral: "10", TokenType: Numeric},
		Token{TokenLiteral: "-20", TokenType: Numeric},
		Token{TokenLiteral: "+", TokenType: Operator},
	}
	testParser(t, results, expected)
}

func TestParseExtraSpace(t *testing.T) {
	inputString := "10 -20 * 5  / 10"
	results := Parse(inputString)
	expected := []Token{
		Token{TokenLiteral: "10", TokenType: Numeric},
		Token{TokenLiteral: "-20", TokenType: Numeric},
		Token{TokenLiteral: "*", TokenType: Operator},
		Token{TokenLiteral: "5", TokenType: Numeric},
		Token{TokenLiteral: "/", TokenType: Operator},
		Token{TokenLiteral: "10", TokenType: Numeric},
	}
	testParser(t, results, expected)
}

func testParser(t *testing.T, results []Token, expected []Token) {
	for i := range results {
		if results[i].TokenLiteral != expected[i].TokenLiteral {
			t.Errorf("%s is expected, but got=%s", expected[i].TokenLiteral, results[i].TokenLiteral)
		}
		if results[i].TokenType != expected[i].TokenType {
			t.Errorf("%s is expected, but got=%s", expected[i].TokenType, results[i].TokenType)
		}
	}

}
