package parser

import (
	"reflect"
)

type ReturnStmt struct {
	Exp Exp
}

func NewReturnStmt(exp Exp) *ReturnStmt {
	return &ReturnStmt{Exp: exp}
}

func (RS *ReturnStmt) Equals(other interface{}) bool {
	castOther := reflect.ValueOf(other)
	otherReturnStmt, ok := reflect.Indirect(castOther).Interface().(ReturnStmt)
	return ok && reflect.TypeOf(other) == reflect.TypeOf(RS) && RS.Exp.Equals(otherReturnStmt)
}
