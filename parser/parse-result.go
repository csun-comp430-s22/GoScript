package parser

type ParseResult[T any] struct {
	Result   T
	Position int
}

func NewParseResult[T any](result T, position int) *ParseResult[T] {
	return &ParseResult[T]{Result: result, Position: position}
}
