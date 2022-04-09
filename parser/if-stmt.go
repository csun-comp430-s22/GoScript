package parser

import (
	"fmt"
	"reflect"
)

type IfStmt struct {
	Guard       Exp
	TrueBranch  Stmt
	FalseBranch Stmt
}

func NewIfStmtOp(guard Exp, trueBranch Stmt, falseBranch Stmt) *IfStmt {
	return &IfStmt{Guard: guard, TrueBranch: trueBranch, FalseBranch: falseBranch}
}

func (IS *IfStmt) Equals(other interface{}) bool {
	if reflect.TypeOf(other) == reflect.TypeOf(IS) {
		// this is a complicated way to cast because of using pointers
		// alternatively can get rid of all pointers in program, pass everything by value
		castOther := reflect.ValueOf(other)
		otherIfStmt, ok := reflect.Indirect(castOther).Interface().(IfStmt)
		return ok && otherIfStmt.Guard.Equals(IS.Guard) && otherIfStmt.TrueBranch.Equals(IS.TrueBranch) && otherIfStmt.FalseBranch.Equals(IS.FalseBranch)
	}

	return false
}

func (IS *IfStmt) String() string {
	return fmt.Sprintf("IfStmt(%s, %s, %s)", IS.Guard, IS.TrueBranch, IS.FalseBranch)
}

func (is *IfStmt) stmt() {}
