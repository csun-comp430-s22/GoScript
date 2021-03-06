package token

import "reflect"

type ElseToken struct {
}

func (et *ElseToken) Equals(other interface{}) bool {
	return reflect.TypeOf(other) == reflect.TypeOf(et)
}

func (et *ElseToken) String() string {
	return "else"
}
