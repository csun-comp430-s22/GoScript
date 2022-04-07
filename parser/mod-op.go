package parser

import "reflect"

type ModOp struct {
}

func (mo *ModOp) Equals(other interface{}) bool {
	return reflect.TypeOf(other) == reflect.TypeOf(mo)
}
