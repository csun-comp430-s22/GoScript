package token

import "reflect"

type OrToken struct {
}

func (ot *OrToken) Equals(other interface{}) bool {
	return reflect.TypeOf(other) == reflect.TypeOf(ot)
}

func (ot *OrToken) String() string {
	return "||"
}
