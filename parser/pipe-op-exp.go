package parser

import (
	"reflect"
)

type PipeOpExp struct {
	Left  Exp
	Op    *PipeOp
	Right Exp
}

func NewPipeOpExp(left Exp, op *PipeOp, right Exp) *PipeOpExp {
	return &PipeOpExp{Left: left, Op: op, Right: right}
}

func (oe *PipeOpExp) Equals(other interface{}) bool {
	if reflect.TypeOf(other) == reflect.TypeOf(oe) {
		// this is a complicated way to cast because of using pointers
		// alternatively can get rid of all pointers in program, pass everything by value
		castOther := reflect.ValueOf(other)
		otherOpExp, ok := reflect.Indirect(castOther).Interface().(PipeOpExp)
		return ok && otherOpExp.Left.Equals(oe.Left) && otherOpExp.Op.Equals(oe.Op) && otherOpExp.Right.Equals(oe.Right)
	}

	return false

}

func (oe *PipeOpExp) exp() {}
