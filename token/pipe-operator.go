package token

import "reflect"

type PipeOperatorToken struct {
}

func (pot *PipeOperatorToken) Equals(other interface{}) bool {
	return reflect.TypeOf(other) == reflect.TypeOf(pot)
}

func (pot *PipeOperatorToken) String() string {
	return "|>"
}
