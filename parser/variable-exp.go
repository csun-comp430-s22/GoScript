package parser

import (
	"reflect"
)

type VariableExp struct {
	Variable Variable
}

func NewVariableExp(variable Variable) *VariableExp {
	return &VariableExp{Variable: variable}
}

func (VE *VariableExp) Equals(other interface{}) bool {
	castOther := reflect.ValueOf(other)
	otherVariableExp, ok := reflect.Indirect(castOther).Interface().(*VariableExp)
	return ok && reflect.TypeOf(other) == reflect.TypeOf(VE) && otherVariableExp.Variable.Equals(VE.Variable)
}

func (ve *VariableExp) exp() {}
