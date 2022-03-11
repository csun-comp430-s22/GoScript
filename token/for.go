package token

import "reflect"

type ForToken struct {
}

func (frt *ForToken) Equals(other interface{}) bool {
	return reflect.TypeOf(other) == reflect.TypeOf(frt)
}

func (frt *ForToken) String() string {
	return "for"
}
