package token

import "reflect"

type ArrowToken struct {
}

func (awt *ArrowToken) Equals(other interface{}) bool {
	return reflect.TypeOf(other) == reflect.TypeOf(awt)
}

func (awt *ArrowToken) String() string {
	return "->"
}
