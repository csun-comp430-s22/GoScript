package token

import (
	"fmt"
	"strings"
	"unicode"
)

type Tokenizer struct {
}

func (t *Tokenizer) Tokenize(input string) []Token {

	tokens := []Token{}
	offset := 0

	for offset < len(input) {
		for unicode.IsSpace(rune(input[offset])) {
			offset++
		}
		if strings.HasPrefix(input[offset:], "true") {
			tokens = append(tokens, &TrueToken{})
			offset += 4
		} else if strings.HasPrefix(input[offset:], "false") {
			tokens = append(tokens, &FalseToken{})
			offset += 5
		} else {
			err := fmt.Errorf("token error: \"%s\" is not a valid token", input[offset:])
			panic(err)
		}
	}

	return tokens
}
