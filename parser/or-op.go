package parser

import "reflect"

type OrOp struct {
}

func (or *OrOp) Equals(other any) bool {
	return reflect.TypeOf(other) == reflect.TypeOf(or)
}
