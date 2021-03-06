package token

import "reflect"

type TrueToken struct {
}

func (tt *TrueToken) Equals(other interface{}) bool {
	return reflect.TypeOf(other) == reflect.TypeOf(tt)
}

func (tt *TrueToken) String() string {
	return "true"
}
