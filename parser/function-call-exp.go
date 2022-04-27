package parser

import (
	"reflect"
)

type FunctionCallExp struct {
	FunctionName *FunctionName
	Params       []Exp
}

func NewFunctionCallExp(functionName *FunctionName, params []Exp) *FunctionCallExp {
	return &FunctionCallExp{FunctionName: functionName, Params: params}
}

func (FCE *FunctionCallExp) Equals(other interface{}) bool {
	if reflect.TypeOf(other) != reflect.TypeOf(FCE) {
		castOther := reflect.ValueOf(other)
		asFunction, ok := reflect.Indirect(castOther).Interface().(FunctionCallExp)
		for i, param := range FCE.Params {
			if !param.Equals(asFunction.Params[i]) {
				return false
			}
		}
		return ok && FCE.FunctionName.Equals(asFunction.FunctionName)
	}
	return false
}

func (fce *FunctionCallExp) exp() {}
