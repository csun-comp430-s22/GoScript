package main

import (
	"fmt"

	"github.com/vSterlin/goscript/token"
)

func main() {

	tokenizer := token.NewTokenizer("true false true")
	tokens := tokenizer.Tokenize()

	for _, t := range tokens {
		fmt.Println(t)
	}

}
