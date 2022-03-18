package main

import (
	"github.com/vSterlin/goscript/parser"
)

// func main() {

// 	tokenizer := token.NewTokenizer("true false if ( ) { } else 100 var123 + - * % / || |> ^ const . && float = string print \"this is a String\"")
// 	tokens := tokenizer.Tokenize()

// 	for _, t := range tokens {
// 		fmt.Println(t)
// 	}

// }

func main() {

	first := parser.NewOpExp(&parser.NumberExp{Number: 1}, &parser.PlusOp{}, &parser.NumberExp{Number: 1})
	second := parser.NewOpExp(&parser.NumberExp{Number: 1}, &parser.PlusOp{}, &parser.NumberExp{Number: 1})

	// fmt.Printf("%#v\n %#v\n", first, second)

	first.Equals(second)

}
