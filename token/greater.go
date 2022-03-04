package token

import "reflect"

type GreaterToken struct {
}

func (gt *GreaterToken) Equals(other interface{}) bool {
	return reflect.TypeOf(other) == reflect.TypeOf(gt)
}

func (gt *GreaterToken) String() string {
	return ">"
}
