package token

import "reflect"

type FnToken struct {
}

func (fnt *FnToken) Equals(other interface{}) bool {
	return reflect.TypeOf(other) == reflect.TypeOf(fnt)
}

func (fnt *FnToken) String() string {
	return "fn"
}
