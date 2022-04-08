package parser

import "reflect"

type AndOp struct {
}

func (and *AndOp) Equals(other any) bool {
	return reflect.TypeOf(other) == reflect.TypeOf(and)
}

type OrOp struct {
}

func (or *OrOp) Equals(other any) bool {
	return reflect.TypeOf(other) == reflect.TypeOf(or)
}

type NegateOp struct {
}

func (neo *NegateOp) Equals(other any) bool {
	return reflect.TypeOf(other) == reflect.TypeOf(neo)
}
