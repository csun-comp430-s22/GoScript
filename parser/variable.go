package parser

import (
	"reflect"
)

type Variable struct {
	Name string
}

func NewVariable(name string) *Variable {
	return &Variable{Name: name}
}

func (v *Variable) Equals(other interface{}) bool {
	castOther := reflect.ValueOf(other)
	otherVariable, ok := reflect.Indirect(castOther).Interface().(Variable)
	return ok && reflect.TypeOf(other) == reflect.TypeOf(v) && otherVariable.Name == v.Name
}
