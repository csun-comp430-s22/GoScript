package parser

import "reflect"

type MinusOp struct {
}

func (mst *MinusOp) Equals(other interface{}) bool {
	return reflect.TypeOf(other) == reflect.TypeOf(mst)
}
