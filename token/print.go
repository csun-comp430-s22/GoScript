package token

import "reflect"

type PrintToken struct {
	name string
}

func (pt *PrintToken) Equals(other interface{}) bool {
	return reflect.TypeOf(other) == reflect.TypeOf(pt)
}

func (pt *PrintToken) String() string {
	return "print"
}
