package token

import (
	"fmt"
	"strings"
	"unicode"
)

type Tokenizer struct {
	tokenMap map[string]Token
	input    []byte
	offset   int
}

func NewTokenizer(input string) *Tokenizer {

	tokenMap := map[string]Token{
		"true":  &TrueToken{},
		"false": &FalseToken{},
		"if":    &IfToken{},
		"(":     &LeftParenToken{},
		")":     &RightParenToken{},
	}

	return &Tokenizer{
		tokenMap: tokenMap,
		input:    []byte(input),
		offset:   0,
	}
}

func (t *Tokenizer) Tokenize() []Token {

	tokens := []Token{}

	for t.offset < len(t.input) {
		for unicode.IsSpace(rune(t.input[t.offset])) {
			t.offset++
		}
		if strings.HasPrefix(string(t.input[t.offset:]), "true") {
			tokens = append(tokens, &TrueToken{})
			t.offset += 4
		} else if strings.HasPrefix(string(t.input[t.offset:]), "false") {
			tokens = append(tokens, &FalseToken{})
			t.offset += 5
		} else {
			err := fmt.Errorf("token error: \"%s\" is not a valid token", t.input[t.offset:])
			panic(err)
		}
	}

	return tokens
}
