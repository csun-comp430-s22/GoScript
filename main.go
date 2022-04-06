package main

import (
	"fmt"

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

	// fmt.Printf("%#v\n %#v\n", first, second)

	one := &parser.ParseResult[parser.Exp]{Result: &parser.NumberExp{1}}
	two := &parser.ParseResult[parser.Exp]{Result: &parser.NumberExp{1}}

	fmt.Printf("one.Equals(two): %v\n", one.Equals(two))

}
