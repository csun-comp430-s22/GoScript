package parser

import (
	"reflect"
)

type Variable struct {
	Name    string
	Mutable bool
}

func NewVariable(name string, mutable bool) *Variable {
	return &Variable{Name: name, Mutable: mutable}
}

func (v *Variable) Equals(other interface{}) bool {
	castOther := reflect.ValueOf(other)
	otherVariable, ok := reflect.Indirect(castOther).Interface().(Variable)
	return ok && reflect.TypeOf(other) == reflect.TypeOf(v) && otherVariable.Name == v.Name && otherVariable.Mutable == v.Mutable
}
