package parser

import "reflect"

type AssignmentOp struct {
}

func (ao *AssignmentOp) Equals(other interface{}) bool {
	return reflect.TypeOf(other) == reflect.TypeOf(ao)
}
