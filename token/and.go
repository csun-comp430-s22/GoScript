package token

import "reflect"

type AndToken struct {
}

func (at *AndToken) Equals(other interface{}) bool {
	return reflect.TypeOf(other) == reflect.TypeOf(at)
}

func (at *AndToken) String() string {
	return "&&"
}
