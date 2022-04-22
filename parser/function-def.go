package parser

import (
	"reflect"
)

type FunctionDef struct {
	Name       FunctionName
	Args       []*Vardec
	Body       Stmt
	ReturnType Type
}

func NewFunctionDef(name FunctionName, args []*Vardec, body Stmt, returnType Type) *FunctionDef {
	return &FunctionDef{Name: name, Args: args, Body: body, ReturnType: returnType}
}

func (FD *FunctionDef) Equals(other interface{}) bool {
	if reflect.TypeOf(other) == reflect.TypeOf(FD) {
		castOther := reflect.ValueOf(other)
		otherFunctionDef, ok := reflect.Indirect(castOther).Interface().(FunctionDef)
		return ok && otherFunctionDef.Name.Equals(FD.Name) && otherFunctionDef.ReturnType.Equals(FD.ReturnType) && equalVardecs(otherFunctionDef.Args, FD.Args) && FD.Body.Equals(otherFunctionDef.Body)
	}
	return false
}

func equalVardecs(a []*Vardec, b []*Vardec) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if !a[i].Equals(b[i]) {
			return false
		}
	}
	return true
}

func (fd *FunctionDef) stmt() {}
