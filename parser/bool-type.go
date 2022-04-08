package parser

import "reflect"

type BoolToken struct {
}

func (bt *BoolToken) Equals(other interface{}) bool {
	return reflect.TypeOf(other) == reflect.TypeOf(bt)
}

func (bt *BoolToken) String() string {
	return "Bool"
}
