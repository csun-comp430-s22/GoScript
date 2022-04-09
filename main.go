package main

import (
	"fmt"

	"github.com/vSterlin/goscript/parser"
	"github.com/vSterlin/goscript/token"
)

// func main() {

// 	tokenizer := token.NewTokenizer("true false if ( ) { } else 100 var123 + - * % / || |> ^ const . && float = string print \"this is a String\"")
// 	tokens := tokenizer.Tokenize()

// 	for _, t := range tokens {
// 		fmt.Println(t)
// 	}

// }

func main() {

	// fmt.Printf("%#v\n %#v\n", first, second)

	// one := &parser.ParseResult[parser.Exp]{Result: &parser.NumberExp{1}}
	// two := &parser.ParseResult[parser.Exp]{Result: &parser.NumberExp{1}}

	// fmt.Printf("one.Equals(two): %v\n", one.Equals(two))

	// tokens := []token.Token{&token.NumberToken{1}, &token.AndToken{}, &token.NumberToken{1}}
	// p := parser.NewParser(tokens)

	// parseResult, _ := p.ParseLogicalExp(0)

	// exp := parser.NewParseResult[parser.Exp](parser.NewOpExp(&parser.NumberExp{1}, &parser.AndOp{}, &parser.NumberExp{1}), 3)

	// fmt.Printf("parseResult: %v\n", parseResult)
	// fmt.Printf("exp: %v\n", exp)

	// fmt.Printf("parseResult.Equals(exp): %v\n", parseResult.Equals(exp))

	tokens := []token.Token{&token.IfToken{}, &token.LeftParenToken{}, &token.NumberToken{1}, &token.EqualsToken{}, &token.NumberToken{1}, &token.RightParenToken{}, &token.LeftCurlyToken{}, &token.NumberToken{1}, &token.RightCurlyToken{}}
	p := parser.NewParser(tokens)

	parseResult, _ := p.ParseStmt(0)

	fmt.Printf("parseResult: %v\n", parseResult)
}
