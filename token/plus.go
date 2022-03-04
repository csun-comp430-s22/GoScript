package token

import "reflect"

type PlusToken struct {
}

func (pt *PlusToken) Equals(other interface{}) bool {
	return reflect.TypeOf(other) == reflect.TypeOf(pt)
}

func (pt *PlusToken) String() string {
	return "+"
}
