package token

import "reflect"

type EqualsToken struct {
}

func (et *EqualsToken) Equals(other interface{}) bool {
	return reflect.TypeOf(other) == reflect.TypeOf(et)
}

func (et *EqualsToken) String() string {
	return "=="
}
