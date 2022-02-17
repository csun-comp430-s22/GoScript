package main

import (
	"fmt"

	"github.com/vSterlin/goscript/token"
)

func main() {

	tokenizer := token.NewTokenizer("true false if ( ) { } else")
	tokens := tokenizer.Tokenize()

	for _, t := range tokens {
		fmt.Println(t)
	}

}
