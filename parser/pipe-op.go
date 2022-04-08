package parser

import "reflect"

type PipeOp struct {
}

func (po *PipeOp) Equals(other any) bool {
	return reflect.TypeOf(other) == reflect.TypeOf(po)
}
