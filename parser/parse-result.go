package parser

import (
	"fmt"
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

	// one will be of type ParseResult[Exp]
	// another will be of type ParseResult[OperatorExp]
	fmt.Printf("reflect.TypeOf(other): %v\n", reflect.TypeOf(other))
	fmt.Printf("reflect.TypeOf(pr): %v\n", reflect.TypeOf(pr))

	// will be false cause of generics
	if reflect.TypeOf(other) == reflect.TypeOf(pr) {
		castOther := reflect.ValueOf(other)
		otherParseResult, ok := reflect.Indirect(castOther).Interface().(ParseResult[T])
		fmt.Printf("castOther: %v\n", castOther)
		fmt.Printf("otherParseResult: %v\n", otherParseResult)
		return ok && pr.Result.Equals(otherParseResult.Result) && pr.Position == otherParseResult.Position
	}
	return false
}
