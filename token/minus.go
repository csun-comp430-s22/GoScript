package token

import "reflect"

type MinusToken struct {
}

func (mst *MinusToken) Equals(other interface{}) bool {
	return reflect.TypeOf(other) == reflect.TypeOf(mst)
}

func (mst *MinusToken) String() string {
	return "-"
}
