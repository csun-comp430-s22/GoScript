package parser

import "reflect"

type ParseResult[T Node] struct {
	Result   T
	Position int
}

func NewParseResult[T Node](result T, position int) *ParseResult[T] {
	return &ParseResult[T]{Result: result, Position: position}
}

func (pr *ParseResult[T]) Equals(other interface{}) bool {

	if reflect.TypeOf(other) == reflect.TypeOf(pr) {
		castOther := reflect.ValueOf(other)
		otherParseResult, ok := reflect.Indirect(castOther).Interface().(ParseResult[T])
		return ok && pr.Result.Equals(otherParseResult.Result) && pr.Position == otherParseResult.Position

	}
	return false
}
