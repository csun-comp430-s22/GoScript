package token

import (
	"fmt"
	"unicode"
)

type Tokenizer struct {
	tokenMap map[string]Token
	input    []byte
	inputPos int
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
		inputPos: 0,
	}
}

func (t *Tokenizer) IsTokenString(input string) bool {
	// if string is in map returns true
	_, ok := t.tokenMap[input]
	return ok
}

func (t *Tokenizer) skipWhitespace() {
	for t.inputPos < len(t.input) && unicode.IsSpace(rune(t.input[t.inputPos])) {
		t.inputPos++
	}
}

func (t *Tokenizer) prefixCharsEqual(probe string) bool {
	targetPos := t.inputPos
	probePos := 0

	for targetPos < len(t.input) && probePos < len(probe) && probe[probePos] == t.input[targetPos] {
		probePos++
		targetPos++
	}
	return probePos == len(probe)
}

func (t *Tokenizer) tryTokenizeOther() Token {
	for key, token := range t.tokenMap {
		if t.prefixCharsEqual(key) {
			t.inputPos += len(key)
			return token
		}
	}
	return nil
}

func (t *Tokenizer) TokenizeSingle() (Token, error) {
	t.skipWhitespace()
	var otherToken Token

	if t.inputPos >= len(t.input) {
		return nil, nil
	} else if otherToken = t.tryTokenizeOther(); otherToken != nil {
		return otherToken, nil
	} else {
		return nil, fmt.Errorf("invalid character %c at position %d", t.input[t.inputPos], t.inputPos)
	}

}

func (t *Tokenizer) Tokenize() []Token {
	tokens := []Token{}
	var current Token

	// second returned value is an error
	for current, _ = t.TokenizeSingle(); current != nil; current, _ = t.TokenizeSingle() {
		tokens = append(tokens, current)
	}

	return tokens
}

// func (t *Tokenizer) Tokenize() []Token {

// 	tokens := []Token{}

// 	if strings.HasPrefix(string(t.input[t.inputPos:]), "true") {
// 		tokens = append(tokens, &TrueToken{})
// 		t.inputPos += 4
// 	} else if strings.HasPrefix(string(t.input[t.inputPos:]), "false") {
// 		tokens = append(tokens, &FalseToken{})
// 		t.inputPos += 5
// 	} else {
// 		err := fmt.Errorf("token error: \"%s\" is not a valid token", t.input[t.inputPos:])
// 		panic(err)
// 	}

// 	return tokens
// }
