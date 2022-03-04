package token

import "reflect"

type RightParenToken struct {
}

func (rpt *RightParenToken) Equals(other interface{}) bool {
	return reflect.TypeOf(other) == reflect.TypeOf(rpt)
}

func (rpt *RightParenToken) String() string {
	return ")"
}
