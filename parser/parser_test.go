package parser

import "testing"

func TestParseSimpleToken(t *testing.T) {
	stringExample := "1 1 +"
	results := Parse(stringExample)
	expected := []token{
		token{tokenLiteral: "1", tokenType: numeric},
		token{tokenLiteral: "1", tokenType: numeric},
		token{tokenLiteral: "+", tokenType: operator},
	}
	testParser(t, results, expected)

}

func TestParseNegativeaInteger(t *testing.T) {
	inputString := "10 -20 +"
	results := Parse(inputString)
	expected := []token{
		token{tokenLiteral: "10", tokenType: numeric},
		token{tokenLiteral: "-20", tokenType: numeric},
		token{tokenLiteral: "+", tokenType: operator},
	}
	testParser(t, results, expected)
}

func TestParseExtraSpace(t *testing.T) {
	inputString := "10 -20 * 5  / 10"
	results := Parse(inputString)
	expected := []token{
		token{tokenLiteral: "10", tokenType: numeric},
		token{tokenLiteral: "-20", tokenType: numeric},
		token{tokenLiteral: "*", tokenType: operator},
		token{tokenLiteral: "5", tokenType: numeric},
		token{tokenLiteral: "/", tokenType: operator},
		token{tokenLiteral: "10", tokenType: numeric},
	}
	testParser(t, results, expected)
}

func testParser(t *testing.T, results []token, expected []token) {
	for i := range results {
		if results[i].tokenLiteral != expected[i].tokenLiteral {
			t.Errorf("%s is expected, but got=%s", expected[i].tokenLiteral, results[i].tokenLiteral)
		}
		if results[i].tokenType != expected[i].tokenType {
			t.Errorf("%s is expected, but got=%s", expected[i].tokenType, results[i].tokenType)
		}
	}

}
