package token

import "reflect"

type MinusToken struct {
}

func (mt *MinusToken) Equals(other interface{}) bool {
	return reflect.TypeOf(other) == reflect.TypeOf(mt)
}

func (mt *MinusToken) String() string {
	return "-"
}
