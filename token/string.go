package token

import "reflect"

type StringToken struct {
}

func (strt *StringToken) Equals(other interface{}) bool {
	return reflect.TypeOf(other) == reflect.TypeOf(strt)
}

func (strt *StringToken) String() string {
	return "string"
}
