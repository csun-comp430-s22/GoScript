package token

import "reflect"

type PowerToken struct {
}

func (pwt *PowerToken) Equals(other interface{}) bool {
	return reflect.TypeOf(other) == reflect.TypeOf(pwt)
}

func (pwt *PowerToken) String() string {
	return "^"
}
