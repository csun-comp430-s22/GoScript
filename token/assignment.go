package token

import "reflect"

type AssignmentToken struct {
}

func (ast *AssignmentToken) Equals(other interface{}) bool {
	return reflect.TypeOf(other) == reflect.TypeOf(ast)
}

func (ast *AssignmentToken) String() string {
	return "="
}
