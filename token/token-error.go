package token

import (
	"fmt"

	"github.com/fatih/color"
)

type TokenizerError struct {
	Message string
}

func (te *TokenizerError) Error() string {
	return te.Message
}

func NewTokenizerError(message string) *TokenizerError {

	return &TokenizerError{Message: message}
}

func handleTokenizerError() {
	if r := recover(); r != nil {
		color.Set(color.FgRed)
		fmt.Println(r)
		color.Unset()
	}
}
