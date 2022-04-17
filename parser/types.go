package parser

import "reflect"

type IntType struct {
}

func (it *IntType) Equals(other any) bool {
	return reflect.TypeOf(other) == reflect.TypeOf(it)

}
