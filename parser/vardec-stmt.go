package parser

import (
	"reflect"
)

type VardecStmt struct {
	Vardec *Vardec
	Exp    Exp
}

func NewVardecStmt(vardec *Vardec, exp Exp) *VardecStmt {
	return &VardecStmt{Vardec: vardec, Exp: exp}
}

func (VS *VardecStmt) Equals(other interface{}) bool {
	if reflect.TypeOf(other) == reflect.TypeOf(VS) {
		castOther := reflect.ValueOf(other)
		otherVardecStmt, ok := reflect.Indirect(castOther).Interface().(VardecStmt)
		return ok && otherVardecStmt.Vardec.Equals(VS.Vardec) && otherVardecStmt.Exp.Equals(VS.Exp)
	}
	return false
}

func (vs *VardecStmt) stmt() {}
