package main

import (
	"fmt"

	"github.com/vSterlin/goscript/token"
)

func main() {

	// tokenizer := token.NewTokenizer("true false if ( ) { } else 100 var123 + - * % / || |> ^ const . && float = string print \"this is a String\"")
	tokenizer := token.NewTokenizer("100")

	tokens := tokenizer.Tokenize()
	numToken := tokens[0]

	anotherToken := &token.NumberToken{Number: 100}

	fmt.Printf("numToken: %#v\n", numToken) // value 100
	fmt.Printf("anotherToken: %v\n", anotherToken)
	isEqual := numToken.Equals(anotherToken)

	fmt.Printf("isEqual: %v\n", isEqual)
	for _, t := range tokens {
		fmt.Println(t)
	}

}
