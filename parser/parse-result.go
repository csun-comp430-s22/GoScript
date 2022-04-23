package parser

import (
	"reflect"

	"github.com/vSterlin/goscript/utils"
)

type ParseResult[T Node] struct {
	Result   T
	Position int
}

func NewParseResult[T Node](result T, position int) *ParseResult[T] {
	return &ParseResult[T]{Result: result, Position: position}
}

func (pr *ParseResult[T]) Equals(other any) bool {

	if utils.GenericTypeCompare(pr, other) {
		// if reflect.TypeOf(other) == reflect.TypeOf(pr) {
		castOther := reflect.ValueOf(other)
		otherParseResult, ok := reflect.Indirect(castOther).Interface().(ParseResult[T])

		return ok && pr.Result.Equals(otherParseResult.Result) && pr.Position == otherParseResult.Position
	}
	return false
}
