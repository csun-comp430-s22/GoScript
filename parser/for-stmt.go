package parser

import (
	"fmt"
	"reflect"
)

type ForStmt struct {
	Guard Exp
	Body  Stmt
}

func NewForStmt(guard Exp, body Stmt) *ForStmt {
	return &ForStmt{Guard: guard, Body: body}
}

func (FS *ForStmt) Equals(other interface{}) bool {
	if reflect.TypeOf(other) == reflect.TypeOf(FS) {
		castOther := reflect.ValueOf(other)
		otherForStmt, ok := reflect.Indirect(castOther).Interface().(ForStmt)
		return ok && otherForStmt.Guard.Equals(FS.Guard) && otherForStmt.Body.Equals(FS.Body)
	}
	return false
}

func (FS *ForStmt) String() string {
	return fmt.Sprintf("ForStmt(%s, %s)", FS.Guard, FS.Body)
}
