package token

import "reflect"

type ConstToken struct {
}

func (ct *ConstToken) Equals(other interface{}) bool {
	return reflect.TypeOf(other) == reflect.TypeOf(ct)
}

func (ct *ConstToken) String() string {
	return "const"
}
