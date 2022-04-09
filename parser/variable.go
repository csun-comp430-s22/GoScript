package parser

import (
	"reflect"
)

type Variable struct {
	Name string
}

func Newvariable(name string) *Variable {
	return &Variable{Name: name}
}

func (V *Variable) Equals(other interface{}) bool {
	castOther := reflect.ValueOf(other)
	otherVariable, ok := reflect.Indirect(castOther).Interface().(*Variable)
	return ok && reflect.TypeOf(other) == reflect.TypeOf(V) && otherVariable.Name == V.Name
}

// func (V *Variable) String() string {
// 	return fmt.Sprintf("Variable(%s)", V.Name)
// }
