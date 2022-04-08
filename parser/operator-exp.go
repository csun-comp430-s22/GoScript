package parser

import (
	"fmt"
	"reflect"
)

type OperatorExp struct {
	Left  Exp
	Op    Operator
	Right Exp
}

func NewOpExp(left Exp, op Operator, right Exp) *OperatorExp {
	return &OperatorExp{Left: left, Op: op, Right: right}
}

func (oe *OperatorExp) Equals(other interface{}) bool {
	if reflect.TypeOf(other) == reflect.TypeOf(oe) {
		// this is a complicated way to cast because of using pointers
		// alternatively can get rid of all pointers in program, pass everything by value
		castOther := reflect.ValueOf(other)
		otherOpExp, ok := reflect.Indirect(castOther).Interface().(OperatorExp)
		return ok && otherOpExp.Left.Equals(oe.Left) && otherOpExp.Op.Equals(oe.Op) && otherOpExp.Right.Equals(oe.Right)
	}

	return false

}

func (oe *OperatorExp) String() string {
	return fmt.Sprintf("OperatorExp(%s, %s, %s)", oe.Left, oe.Op, oe.Right)
}

func (oe *OperatorExp) exp() {}
