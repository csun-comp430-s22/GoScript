package parser

import "reflect"

type IntType struct {
}

func (it *IntType) Equals(other any) bool {
	return reflect.TypeOf(other) == reflect.TypeOf(it)
}

type BoolType struct {
}

func (bt *BoolType) Equals(other any) bool {
	return reflect.TypeOf(other) == reflect.TypeOf(bt)
}
