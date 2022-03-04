package token

import "reflect"

type FloatToken struct {
}

func (flt *FloatToken) Equals(other interface{}) bool {
	return reflect.TypeOf(other) == reflect.TypeOf(flt)
}

func (flt *FloatToken) String() string {
	return "float"
}
