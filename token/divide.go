package token

import "reflect"

type DivToken struct {
}

func (dt *DivToken) Equals(other interface{}) bool {
	return reflect.TypeOf(other) == reflect.TypeOf(dt)
}

func (dt *DivToken) String() string {
	return "/"
}
