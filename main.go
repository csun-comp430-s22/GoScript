package main

import (
	"fmt"

	"github.com/vSterlin/goscript/token"
)

func main() {

	tokenizer := token.Tokenizer{}
	tokens := tokenizer.Tokenize("true false true")

	for _, t := range tokens {
		fmt.Println(t)
	}

}
