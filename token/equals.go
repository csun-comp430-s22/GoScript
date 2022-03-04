package token

import "reflect"

type EqualsToken struct {
}

func (eqt *EqualsToken) Equals(other interface{}) bool {
	return reflect.TypeOf(other) == reflect.TypeOf(eqt)
}

func (eqt *EqualsToken) String() string {
	return "="
}
