package main

import (
	"fmt"

	"github.com/vSterlin/goscript/token"
)

func main() {

	// fmt.Printf("%#v\n %#v\n", first, second)

	// one := &parser.ParseResult[parser.Exp]{Result: &parser.NumberExp{1}}
	// two := &parser.ParseResult[parser.Exp]{Result: &parser.NumberExp{1}}

	// fmt.Printf("one.Equals(two): %v\n", one.Equals(two))

	// tokens := []token.Token{
	// 	&token.IfToken{},
	// 	&token.LeftParenToken{},
	// 	&token.NumberToken{1},
	// 	&token.EqualsToken{},
	// 	&token.NumberToken{1},
	// 	&token.RightParenToken{},
	// 	&token.LeftCurlyToken{},
	// 	&token.RightCurlyToken{},
	// 	&token.ElseToken{},
	// 	&token.LeftCurlyToken{},
	// 	&token.RightCurlyToken{},
	// }

	// tokens := []token.Token{
	// 	&token.LeftCurlyToken{},
	// 	&token.RightCurlyToken{},
	// }

	// tokens := []token.Token{&token.NumberToken{1}, &token.EqualsToken{}, &token.NumberToken{1}}

	// tokenizer := token.NewTokenizer(`
	// 	fn int funcName(x int){}
	// 	fn int anotherFunc(a int){}
	// 	fn int thirdFunc(a int, b int){}`)

	// tokenizer := token.NewTokenizer("fn int test(x int){}")

	// tokens := tokenizer.Tokenize()

	// for i, t := range tokens {
	// 	fmt.Printf("%d: %#v\n", i, t)
	// }

	// p := parser.NewParser(tokens)
	// parseRes, _ := p.ParseFunctionDefinition(0)

	// expected := parser.NewParseResult(parser.NewFunctionDef(
	// 	parser.NewFunctionName("test"),
	// 	[]*parser.Vardec{parser.NewVardec(parser.NewVariable("x"), &parser.IntType{})},
	// 	parser.NewBlockStmt([]parser.Stmt{}),
	// 	&parser.IntType{},
	// ), len(tokens)-1)

	// fmt.Printf("parseRes.Equals(expected): %v\n", parseRes.Equals(expected))

	tokenizer := token.NewTokenizer("test(1,2)")

	tokens := tokenizer.Tokenize()

	for i, t := range tokens {
		fmt.Printf("%d: %#v\n", i, t)
	}

	// p := parser.NewParser(tokens)

	// parseResult, _ := p.ParseFunctionCall(0)

	// expected := parser.NewParseResult(parser.NewFunctionCallExp(parser.NewFunctionName("test"), []parser.Exp{&parser.IntLiteralExp{1}, &parser.IntLiteralExp{2}}), 5)

	// if !parseResult.Equals(expected) {
	// 	fmt.Printf("expected: %#v\n", expected.Result)
	// 	fmt.Printf("got: %#v\n", parseResult.Result)
	// 	fmt.Printf("result: %#v\n", parseResult.Result.FunctionName.Equals(expected.Result.FunctionName))

	// }
}
