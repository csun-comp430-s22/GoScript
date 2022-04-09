package parser

import (
	"reflect"
)

type FunctionName struct {
	Name string
}

func NewFunctionName(name string) *FunctionName {
	return &FunctionName{Name: name}
}

func (FN *FunctionName) Equals(other interface{}) bool {
	castOther := reflect.ValueOf(other)
	otherFunctionName, ok := reflect.Indirect(castOther).Interface().(FunctionName)
	return ok && reflect.TypeOf(other) == reflect.TypeOf(FN) && FN.Name == otherFunctionName.Name
}
