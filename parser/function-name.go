package parser

import (
	"fmt"
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

	fmt.Printf("castOther: %#v\n", castOther)
	fmt.Printf("otherFunctionName: %#v\n", otherFunctionName)

	fmt.Printf("reflect.TypeOf(other): %v\n", reflect.TypeOf(other))
	fmt.Printf("reflect.TypeOf(FN): %v\n", reflect.TypeOf(FN))
	return ok && reflect.TypeOf(other) == reflect.TypeOf(FN) && FN.Name == otherFunctionName.Name
}
