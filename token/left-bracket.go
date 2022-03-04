package token

import "reflect"

type LeftBracketToken struct {
}

func (lbt *LeftBracketToken) Equals(other interface{}) bool {
	return reflect.TypeOf(other) == reflect.TypeOf(lbt)
}

func (lbt *LeftBracketToken) String() string {
	return "["
}
