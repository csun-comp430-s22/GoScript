package parser

import "reflect"

type LessOp struct {
}

func (lo *LessOp) Equals(other interface{}) bool {
	return reflect.TypeOf(other) == reflect.TypeOf(lo)
}

type GreaterOp struct {
}

func (gop *GreaterOp) Equals(other interface{}) bool {
	return reflect.TypeOf(other) == reflect.TypeOf(gop)
}

type EqualsOp struct {
}

func (eo *EqualsOp) Equals(other interface{}) bool {
	return reflect.TypeOf(other) == reflect.TypeOf(eo)
}

type NotEqualsOp struct {
}

func (neo *NotEqualsOp) Equals(other interface{}) bool {
	return reflect.TypeOf(other) == reflect.TypeOf(neo)
}

type LessEqualOp struct {
}

func (leo *LessEqualOp) Equals(other interface{}) bool {
	return reflect.TypeOf(other) == reflect.TypeOf(leo)
}

type GreaterEqualOp struct {
}

func (geop *GreaterEqualOp) Equals(other interface{}) bool {
	return reflect.TypeOf(other) == reflect.TypeOf(geop)
}
