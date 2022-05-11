package parser

import (
	"reflect"
)

type ParseResult[T Node] struct {
	Result   T
	Position int
}

func NewParseResult[T Node](result T, position int) *ParseResult[T] {
	return &ParseResult[T]{Result: result, Position: position}
}

func (pr *ParseResult[T]) Equals(other any) bool {

	// deep equals wont't work if generic type is not the same
	// for example some ParseResult[Stmt] will not equal some ParseResult[FunctionDefinition]
	// even though FunctionDefinition implements Stmt interface. Gotta take that into account
	return reflect.DeepEqual(pr, other)
}
