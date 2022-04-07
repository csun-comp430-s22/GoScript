package parser

import "reflect"

type DivideOp struct {
}

func (do *DivideOp) Equals(other interface{}) bool {
	return reflect.TypeOf(other) == reflect.TypeOf(do)
}
