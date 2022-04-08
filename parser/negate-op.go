package parser

import "reflect"

type NegateOp struct {
}

func (neo *NegateOp) Equals(other any) bool {
	return reflect.TypeOf(other) == reflect.TypeOf(neo)
}
