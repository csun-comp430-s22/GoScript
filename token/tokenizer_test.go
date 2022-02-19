package token

import (
	"testing"
)

func TestInvalidInput(t *testing.T) {

	input := "badinput"

	tokenizer := NewTokenizer(input)
	_, err := tokenizer.TokenizeSingle()
	if err == nil {
		t.Error("no error returned")
	}
	if _, ok := err.(*TokenizerError); !ok {
		t.Error("unexpected error returned")
	}

}
