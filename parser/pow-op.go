package parser

import "reflect"

type PowOp struct {
}

func (po *PowOp) Equals(other any) bool {
	return reflect.TypeOf(other) == reflect.TypeOf(po)
}
