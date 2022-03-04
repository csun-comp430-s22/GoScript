package token

import "reflect"

type DotToken struct {
}

func (dot *DotToken) Equals(other interface{}) bool {
	return reflect.TypeOf(other) == reflect.TypeOf(dot)
}

func (dot *DotToken) String() string {
	return "."
}
