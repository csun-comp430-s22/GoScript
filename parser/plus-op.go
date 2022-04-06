package parser

import "reflect"

type PlusOp struct {
}

func (po *PlusOp) Equals(other interface{}) bool {
	return reflect.TypeOf(other) == reflect.TypeOf(po)
}
