package parser

import (
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

func (oe *OperatorExp) Equals(other any) bool {
	return reflect.DeepEqual(oe, other)
}

func (oe *OperatorExp) exp() {}
