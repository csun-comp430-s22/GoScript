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
