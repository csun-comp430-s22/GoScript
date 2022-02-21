package token

import (
	"testing"
)

func TestTokenizerError(t *testing.T) {

	input := "badinput"

	tokenizer := NewTokenizer(input)
	_, err := tokenizer.TokenizeSingle()
	if err == nil {
		t.Error("no error returned")
	}
	/* same as writing
	_, ok := err.(*TokenizerError)
	if !ok {
		t.Error("unexpected error returned")
	}
	*/
	if _, ok := err.(*TokenizerError); !ok {
		t.Error("unexpected error returned")
	}

}

func TestTokenizeIf(t *testing.T) {
	input := "if"

	tokenizer := NewTokenizer(input)
	tokens := tokenizer.Tokenize()
	if !tokens[0].Equals(&IfToken{}) {
		t.Error()
	}

}

func TestAllKeywords(t *testing.T) {
	type keywordTest struct {
		Input          string
		ExpectedResult Token
	}

	keywordTests := []keywordTest{
		{"if", &IfToken{}}}

	for _, test := range keywordTests {
		t.Run(test.Input, func(t *testing.T) {
			tokenizer := NewTokenizer(test.Input)
			tokens := tokenizer.Tokenize()
			actual := tokens[0]
			if !actual.Equals(test.ExpectedResult) {
				t.Errorf("expected token: \"%s\", got: %s", test.ExpectedResult, actual)
			}
		})
	}

}
