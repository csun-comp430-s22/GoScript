package token

import "reflect"

type SemiColonToken struct {
}

func (sct *SemiColonToken) Equals(other interface{}) bool {
	return reflect.TypeOf(other) == reflect.TypeOf(sct)
}

func (sct *SemiColonToken) String() string {
	return ";"
}
