package parser

import "reflect"

type AndOp struct {
}

func (and *AndOp) Equals(other any) bool {
	return reflect.TypeOf(other) == reflect.TypeOf(and)
}
