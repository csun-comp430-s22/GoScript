package token

type TokenizerError struct {
	Message string
}

func (te *TokenizerError) Error() string {
	return te.Message
}

func NewTokenizerError(message string) *TokenizerError {

	return &TokenizerError{Message: message}
}
