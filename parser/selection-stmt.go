package parser

import (
	"reflect"
)

type IfStmt struct {
	Guard  Exp
	Branch Stmt
}

func NewIfStmt(guard Exp, branch Stmt) *IfStmt {
	return &IfStmt{Guard: guard, Branch: branch}
}

func (IS *IfStmt) Equals(other interface{}) bool {
	if reflect.TypeOf(other) == reflect.TypeOf(IS) {
		// this is a complicated way to cast because of using pointers
		// alternatively can get rid of all pointers in program, pass everything by value
		castOther := reflect.ValueOf(other)
		otherIfStmt, ok := reflect.Indirect(castOther).Interface().(IfStmt)
		return ok && otherIfStmt.Guard.Equals(IS.Guard) && otherIfStmt.Branch.Equals(IS.Branch)
	}

	return false
}

func (is *IfStmt) stmt() {}

type IfElseStmt struct {
	Guard       Exp
	TrueBranch  Stmt
	FalseBranch Stmt
}

func NewIfElseStmt(guard Exp, trueBranch Stmt, falseBranch Stmt) *IfElseStmt {
	return &IfElseStmt{Guard: guard, TrueBranch: trueBranch, FalseBranch: falseBranch}
}

func (IS *IfElseStmt) Equals(other interface{}) bool {
	if reflect.TypeOf(other) == reflect.TypeOf(IS) {
		// this is a complicated way to cast because of using pointers
		// alternatively can get rid of all pointers in program, pass everything by value
		castOther := reflect.ValueOf(other)
		otherIfStmt, ok := reflect.Indirect(castOther).Interface().(IfElseStmt)
		return ok && otherIfStmt.Guard.Equals(IS.Guard) && otherIfStmt.TrueBranch.Equals(IS.TrueBranch) && otherIfStmt.FalseBranch.Equals(IS.FalseBranch)
	}

	return false
}

func (is *IfElseStmt) stmt() {}
