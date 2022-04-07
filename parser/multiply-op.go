package parser

import "reflect"

type MultiplyOp struct {
}

func (mo *MultiplyOp) Equals(other interface{}) bool {
	return reflect.TypeOf(other) == reflect.TypeOf(mo)
}
