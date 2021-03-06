package token

import "reflect"

type VariableToken struct {
	Name string
}

func (vt *VariableToken) Equals(other interface{}) bool {
	return reflect.TypeOf(other) == reflect.TypeOf(vt)
}

func (vt *VariableToken) String() string {
	return "Variable(" + vt.Name + ")"
}
