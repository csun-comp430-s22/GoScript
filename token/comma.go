package token

import "reflect"

type CommaToken struct {
}

func (cmt *CommaToken) Equals(other interface{}) bool {
	return reflect.TypeOf(other) == reflect.TypeOf(cmt)
}

func (cmt *CommaToken) String() string {
	return ","
}
