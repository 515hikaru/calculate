package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/515hikaru/calculate/calc"
	"github.com/515hikaru/calculate/parser"
)

func main() {
	var s string
	sc := bufio.NewScanner(os.Stdin)
	if sc.Scan() {
		s = sc.Text()
	} else {
		panic("cannot load from stdinput")
	}
	tokens := parser.Parse(s)
	calc := calc.Calc(tokens)
	fmt.Printf("%v\n", calc)
}
