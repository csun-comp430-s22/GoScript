package token

import "reflect"

type RightBracketToken struct {
}

func (rbt *RightBracketToken) Equals(other interface{}) bool {
	return reflect.TypeOf(other) == reflect.TypeOf(rbt)
}

func (rbt *RightBracketToken) String() string {
	return "]"
}
