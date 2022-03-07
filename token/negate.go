package token

import "reflect"

type NegateToken struct {
}

func (ngt *NegateToken) Equals(other interface{}) bool {
	return reflect.TypeOf(other) == reflect.TypeOf(ngt)
}

func (ngt *NegateToken) String() string {
	return "!"
}
