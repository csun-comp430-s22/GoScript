package parser

type ParserError struct {
	Message string
}

func (pe *ParserError) Error() string {
	return pe.Message
}

func NewParserError(message string) *ParserError {

	return &ParserError{Message: message}
}
