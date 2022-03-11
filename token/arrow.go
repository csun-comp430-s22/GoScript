package token

import "reflect"

type ArrowToken struct {
}

func (art *ArrowToken) Equals(other interface{}) bool {
	return reflect.TypeOf(other) == reflect.TypeOf(art)
}

func (art *ArrowToken) String() string {
	return "->"
}
