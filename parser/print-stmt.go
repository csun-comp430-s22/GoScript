package parser

import (
	"reflect"
)

type PrintStmt struct {
	Exp Exp
}

func NewPrintStmtOp(exp Exp) *PrintStmt {
	return &PrintStmt{Exp: exp}
}

func (PS *PrintStmt) Equals(other interface{}) bool {
	// idk if casting was needed or not
	castOther := reflect.ValueOf(other)
	otherPrintStmt, ok := reflect.Indirect(castOther).Interface().(PrintStmt)
	return ok && reflect.TypeOf(other) == reflect.TypeOf(PS) && PS.Exp.Equals(otherPrintStmt)
}

// func (PS *PrintStmt) String() string {
// 	return fmt.Sprintf("PrintStmt(%s)", PS.Exp)
// }
